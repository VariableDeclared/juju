run_secrets_juju() {
	echo

	file="${TEST_DIR}/model-secrets-juju.txt"

	juju --show-log add-model "model-secrets-juju" --config secret-store=juju
	juju --show-log deploy easyrsa
	wait_for "easyrsa" "$(idle_condition "easyrsa")"
	
	secret_owned_by_easyrsa_0=$(juju exec --unit easyrsa/0 -- secret-add --owner unit owned-by=easyrsa/0 | cut -d: -f2)
	secret_owned_by_easyrsa=$(juju exec --unit easyrsa/0 -- secret-add owned-by=easyrsa | cut -d: -f2)
	
	juju exec --unit easyrsa/0 -- secret-ids | grep "$secret_owned_by_easyrsa"
	juju exec --unit easyrsa/0 -- secret-ids | grep "$secret_owned_by_easyrsa_0"
	
	echo "Set a label for the unit owned secret $secret_owned_by_easyrsa_0."
	juju exec --unit easyrsa/0 -- secret-set "$secret_owned_by_easyrsa_0" --label=easyrsa_0
	echo "Set a consumer label for the app owned secret $secret_owned_by_easyrsa."
	juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa" --label=easyrsa
	
	# secret-get by URI - content.
	juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa_0" | grep 'owned-by: easyrsa/0'
	juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa" | grep 'owned-by: easyrsa'
	
	# secret-get by URI - metadata.
	juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa_0" --metadata --format json | jq ".${secret_owned_by_easyrsa_0}.owner" | grep unit
	juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa" --metadata --format json | jq ".${secret_owned_by_easyrsa}.owner" | grep application
	
	# secret-get by label or consumer label - content.
	juju exec --unit easyrsa/0 -- secret-get --label=easyrsa_0 | grep 'owned-by: easyrsa/0'
	juju exec --unit easyrsa/0 -- secret-get --label=easyrsa | grep 'owned-by: easyrsa'
	
	# secret-get by label - metadata.
	juju exec --unit easyrsa/0 -- secret-get --label=easyrsa_0 --metadata --format json | jq ".${secret_owned_by_easyrsa_0}.label" | grep easyrsa_0
	
	juju --show-log deploy etcd
	wait_for "etcd" "$(idle_condition "etcd")"

	echo "Relate workload in consume model with offer"
	juju --show-log integrate etcd easyrsa
	
	wait_for "easyrsa" '.applications["etcd"] | .relations.certificates[0]'
	
	relation_id=$(juju --show-log show-unit easyrsa/0 --format json | jq '."easyrsa/0"."relation-info"[0]."relation-id"')
	juju exec --unit easyrsa/0 -- secret-grant "$secret_owned_by_easyrsa_0" -r "$relation_id"
	juju exec --unit easyrsa/0 -- secret-grant "$secret_owned_by_easyrsa" -r "$relation_id"
	
	# secret-get by URI - consume content.
	juju exec --unit etcd/0 -- secret-get "$secret_owned_by_easyrsa_0" | grep 'owned-by: easyrsa/0'
	juju exec --unit etcd/0 -- secret-get "$secret_owned_by_easyrsa" | grep 'owned-by: easyrsa'
	
	juju exec --unit etcd/0 -- secret-get "$secret_owned_by_easyrsa_0" --label=consumer_label_secret_owned_by_easyrsa_0 | grep 'owned-by: easyrsa/0'
	juju exec --unit etcd/0 -- secret-get "$secret_owned_by_easyrsa" --label=consumer_label_secret_owned_by_easyrsa | grep 'owned-by: easyrsa'
	
	juju exec --unit etcd/0 -- secret-get --label=consumer_label_secret_owned_by_easyrsa_0 | grep 'owned-by: easyrsa/0'
	juju exec --unit etcd/0 -- secret-get --label=consumer_label_secret_owned_by_easyrsa | grep 'owned-by: easyrsa'
	
	# secret-revoke by relation ID.
	juju exec --unit easyrsa/0 -- secret-revoke $secret_owned_by_easyrsa --relation "$relation_id"
	juju exec --unit etcd/0 -- secret-get "$secret_owned_by_easyrsa" | grep 'permission denied'
	
	# secret-revoke by app name.
	juju exec --unit easyrsa/0 -- secret-revoke $secret_owned_by_easyrsa_0 --app etcd
	juju exec --unit etcd/0 -- secret-get "$secret_owned_by_easyrsa_0" | grep 'permission denied'
	
	# secret-remove.
	juju exec --unit easyrsa/0 -- secret-remove "$secret_owned_by_easyrsa_0"
	juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa_0" | grep 'not found'
	# TODO: enable once we fix: https://bugs.launchpad.net/juju/+bug/1994919.
	# juju exec --unit easyrsa/0 -- secret-remove "$secret_owned_by_easyrsa"
	# juju exec --unit easyrsa/0 -- secret-get "$secret_owned_by_easyrsa" | grep 'not found'

	destroy_model "model-secrets-juju"
}

test_secrets_juju() {
	if [ "$(skip 'test_secrets_juju')" ]; then
		echo "==> TEST SKIPPED: test_secrets_juju"
		return
	fi

	(
		set_verbosity

		cd .. || exit

		run "run_secrets_juju"
	)
}

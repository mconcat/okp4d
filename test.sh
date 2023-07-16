# export MNEMONIC_VALIDATOR="island position immense mom cross enemy grab little deputy tray hungry detect state helmet tomorrow trap expect admit inhale present vault reveal scene atom";
# echo $MNEMONIC_VALIDATOR | okp4d keys add validator --recover

export ADDR=okp41p8u47en82gmzfm259y6z93r9qe63l25dfwwng6

okp4d tx wasm store ./okp4_predicates.wasm --gas 20000000 --from validator --chain-id okp4-localnet --sequence 6
# okp4d tx wasm instantiate 1 '{"name":"okp4_predicates"}' --label "okp4_predicates" --gas 20000000 --no-admin --from validator --chain-id okp4-localnet 
package p2pv1

/*
Validator = entitas yang ikut finality voting
- PublicKey dipakai verifikasi signature
- Power menentukan bobot suara
*/

type Validator struct {
	ID        string
	PublicKey []byte
	Power     uint64
}

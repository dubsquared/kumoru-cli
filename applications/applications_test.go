package applications

import "testing"

func TestReadCertificatesEmpty(t *testing.T) {

	var cert, key, ca string
	expected := ""

	result := readCertificates(&cert, &key, &ca)

	if result != expected {
		t.Errorf("result == %v, want %v", result, expected)
	}
}

//TODO Implement following test case with file reads
/*func TestReadCertificatesExists(t *testing.T) {

	cert := "mycert"
	key := "mykey"
	ca := "myca"

	expected := "{\"certificate\": \"mycert\", \"private_key\": \"mykey\", \"certificate_authority\": \"myca\"}"

	result := readCertificates(&cert, &key, &ca)

	if result != expected {
		t.Errorf("result == %v, want %v", result, expected)
	}
}*/

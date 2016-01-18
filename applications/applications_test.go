/*
Copyright 2016 Kumoru.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

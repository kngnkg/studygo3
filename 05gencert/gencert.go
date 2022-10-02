package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// 開発/テスト用のSSL証明書と秘密鍵を生成する。
//
//	SSL証明書: cert.pem
//	サーバ用の秘密鍵: key.pem
func main() {

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 証明書の構成を設定するための構造体
	// シリアル番号(SerialNumber)：本来は認証局によって発行される一意の番号
	// 識別名(Subject)
	// 有効期間：1年間
	// KeyUsage/ExtKeyUsage：このX.509証明書がサーバ認証に使用されることを示す
	// IPアドレス127.0.0.1でのみ効力を持つ
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	// 秘密鍵の生成
	// 秘密鍵の構造体には公開鍵が含まれる
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// SSL証明書の作成
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 秘密鍵をPEM符号化してファイルに保存
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

}

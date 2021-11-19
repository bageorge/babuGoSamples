package main
 
import (
    "crypto/aes"
    "encoding/hex"
    "fmt"
    "os"
    "log"
)


func main() {
 
    dataFile := os.Args[1]
    actType  := os.Args[2]
    fmt.Println("Input data file => ", dataFile)
    fmt.Println("Action encrypt/decrypt => ", os.Args[2])
    fmt.Println("Input file contents")

    dat, err := os.ReadFile(dataFile)
    CheckError(err)
    fmt.Print(string(dat))

    switch actType {
    case "enc":
         fmt.Println("Reading each entry and sending for encryption")
         //config := ReadConfig()
    case "dec":
         fmt.Println("Reading each entry and sending for decryption")
    }

    // cipher key
    //key := "thisis32bitlongpassphraseimusing"
    
    if actType=="enc" {  
       // plaintext
       //pt := "This is a secret"
       config := ReadConfig()
       pt := config.sysdbuser 
       fmt.Println(pt)

       // Read each key:value from the file and set the value as plain text
 
       //c := EncryptAES([]byte(key), pt)
 
      // plaintext
      fmt.Println(pt)
 
      // ciphertext
      //fmt.Println(c)
    } else if actType=="dec" {
 
       // decrypt
       //DecryptAES([]byte(key), c)
       fmt.Println("Decryption data from file")
   } else {
       fmt.Println("Invalid Action, only enc/dec supported")
   }
}
 
func EncryptAES(key []byte, plaintext string) string {
 
    c, err := aes.NewCipher(key)
    CheckError(err)
 
    out := make([]byte, len(plaintext))
 
    c.Encrypt(out, []byte(plaintext))
 
    return hex.EncodeToString(out)
}
 
func DecryptAES(key []byte, ct string) {
    ciphertext, _ := hex.DecodeString(ct)
 
    c, err := aes.NewCipher(key)
    CheckError(err)
 
    pt := make([]byte, len(ciphertext))
    c.Decrypt(pt, ciphertext)
 
    s := string(pt[:])
    fmt.Println("DECRYPTED:", s)
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

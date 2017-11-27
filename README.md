# cfgcrypt

CfgCrypt or config crypt is a cli tool to encrypt values in a text configuration file for use within a secure application.

## Concept

 The basic idea is that you write your configuration file in whatever text format you prefer. Then you run cfgcrypt on the file to encrypt the variables you want hidden. From there your application base 64 decodes the encrypted text, and runs AES 128 decryption, in CBC mode, with PKCS7 padding, using the key from the generated key file. The iv for the decryption is included in the first set of bytes.

 So that cfgcrypt knows what to encrypt you must give it a prefix and a postfix string that will delimit the values that you wish to encrypt. These values will be encrypted in-place. If an encryption key is not passed to the utility then one will be randomly generated and placed next to the encrypted file with 0600 file permissions and a ".key" file extension pre-pended to the original file's name.

 ## Decryption logic

 Pseudo-code example of decryption process:

 ```
 configData = io.ReadConfigFile(fileName)
 key = io.ReadFileBytes(fileName + ".key")
 decoded = base64.decode(configData.secretValue)
 iv = getFirst16Bytes(decoded)
 encrypted = getBytesAfterFirst16(decoded)
 cipher = encryption.Mode("AES/CBC/PKCS7Padding")
 configData.secretValue = cipher.Decrypt(iv, encrypted, key)
 ```

 ## Usage

 ```
 cfgcrypt [textfile] ...
	textfile	Text file to encrypt. (required)
  -key string
    	Base64 encoded encryption key, if not specified one will be generated
  -postfix string
    	Post string denoting end of value to be encrypted (default "}}#")
  -prefix string
    	Prefix string denoting start of value to be encrypted (default "#{{")
  -force bool
  	Overwrites key file if it exists
```

## Example

Examples are located in the `examples` folder with the original unencrypted files in `examples/original` and the encrypted output in `examples/encrypted`

## How To Build

To build run
```bash
godep save
godep go build
```

## Future development

I'm considering the following upgrades:
- Support for more encryption modes/algorithms
- Safer error messages to avoid leaking any security details (not certain if necessary)

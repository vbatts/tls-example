tls-example
===========

Golang TLS example. x509 certificate create and sign.

Usage:


    go run gen.go
    2013/11/25 14:47:48 write to ca.pem
    2013/11/25 14:47:48 write to ca.key
    2013/11/25 14:47:49 write to cert2.pem
    2013/11/25 14:47:49 write to cert2.key
    2013/11/25 14:47:49 check signature true
    
    go run server.go
    2015/05/28 16:30:57 server: listening
    2015/05/28 16:31:00 server: accepted from 127.0.0.1:49114
    2015/05/28 16:31:00 server: conn: waiting
    2015/05/28 16:31:00 {[US] [Acme] [AcmeClient] [] [] [] []   [{2.5.4.6 US} {2.5.4.10 Acme} {2.5.4.11 AcmeClient}]}
    2015/05/28 16:31:00 server: conn: echo "Hello\n"
    2015/05/28 16:31:00 server: conn: wrote 6 bytes
    2015/05/28 16:31:00 server: conn: waiting
      
    go run client.go
    2015/05/28 16:30:40 client: connected to:  127.0.0.1:11111
    [48 129 159 48 13 6 9 42 134 72 134 247 13 1 1 1 5 0 3 129 141 0 48 129 137 2 129 129 0 227 141 114 184 57 195 237 224 73 161 43 212 189 129 89 94 39 139 12 41 23 76 14 173 89 92 47 175 76 237 70 157 211 24 9 211 86 186 197 74 252 113 78 216 42 219 83 48 38 60 26 129 197 195 139 232 9 94 41 115 194 126 67 159 127 38 23 89 89 220 23 25 85 10 145 232 130 141 225 5 170 155 165 66 34 184 1 224 249 231 81 16 151 188 229 41 128 18 125 203 229 117 17 183 192 248 26 91 214 240 15 216 238 212 21 15 75 210 64 122 227 184 161 191 150 89 123 231 2 3 1 0 1] <nil>
    {[US] [Acme] [AcmeServer] [] [] [] []   [{2.5.4.6 US} {2.5.4.10 Acme} {2.5.4.11 AcmeServer}]}
    2015/05/28 16:30:40 client: handshake:  true
    2015/05/28 16:30:40 client: mutual:  true
    2015/05/28 16:30:40 client: wrote "Hello\n" (6 bytes)
    2015/05/28 16:30:40 client: read "Hello\n" (6 bytes)

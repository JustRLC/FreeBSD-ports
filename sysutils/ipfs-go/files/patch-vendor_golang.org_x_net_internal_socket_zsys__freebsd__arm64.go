--- vendor/golang.org/x/net/internal/socket/zsys_freebsd_arm64.go.orig	2020-04-04 20:13:32 UTC
+++ vendor/golang.org/x/net/internal/socket/zsys_freebsd_arm64.go
@@ -0,0 +1,53 @@
+// Code generated by cmd/cgo -godefs; DO NOT EDIT.
+// cgo -godefs defs_freebsd.go
+
+package socket
+
+type iovec struct {
+	Base *byte
+	Len  uint64
+}
+
+type msghdr struct {
+	Name       *byte
+	Namelen    uint32
+	Pad_cgo_0  [4]byte
+	Iov        *iovec
+	Iovlen     int32
+	Pad_cgo_1  [4]byte
+	Control    *byte
+	Controllen uint32
+	Flags      int32
+}
+
+type cmsghdr struct {
+	Len   uint32
+	Level int32
+	Type  int32
+}
+
+type sockaddrInet struct {
+	Len    uint8
+	Family uint8
+	Port   uint16
+	Addr   [4]byte /* in_addr */
+	Zero   [8]int8
+}
+
+type sockaddrInet6 struct {
+	Len      uint8
+	Family   uint8
+	Port     uint16
+	Flowinfo uint32
+	Addr     [16]byte /* in6_addr */
+	Scope_id uint32
+}
+
+const (
+	sizeofIovec   = 0x10
+	sizeofMsghdr  = 0x30
+	sizeofCmsghdr = 0xc
+
+	sizeofSockaddrInet  = 0x10
+	sizeofSockaddrInet6 = 0x1c
+)

package utils

// Float32ToUint converts float32 to uint (since oapi-codegen uses float32 for IDs)
func Float32ToUint(f float32) uint {
	return uint(f)
}

package appCommon

func MakePath(spaceId, fileName string) string {
	return Join("/", S3PathSpaceImage, spaceId, fileName)
}
func MakeSpaceUrl(spaceId, fileName string) string {
	return Join("/", S3Domain, "space", spaceId, fileName)
}

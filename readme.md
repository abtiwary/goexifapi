# Image Exif API

Todo: work in progress...

#### Test the API using curl for now, for example:

```
curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@./images/193A1268_.JPG" http://localhost:8888/api/v1/exif
```

##### Example response:

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 04 Jul 2021 15:13:26 GMT
Content-Length: 1286

{"ApertureValue":"\"458752/65536\"","Artist":"\"\"","ColorSpace":"1","ComponentsConfiguration":"\"\"","Copyright":"\"\"","CustomRendered":"0","DateTime":"\"2017:06:19 18:41:56\"","DateTimeDigitized":"\"2017:06:11 06:27:16\"","DateTimeOriginal":"\"2017:06:11 06:27:16\"","ExifIFDPointer":"222","ExifVersion":"\"0230\"","ExposureBiasValue":"\"0/1\"","ExposureMode":"1","ExposureProgram":"1","ExposureTime":"\"1/1250\"","FNumber":"\"11/1\"","Flash":"16","FlashpixVersion":"\"0100\"","FocalLength":"\"29/1\"","FocalPlaneResolutionUnit":"2","FocalPlaneXResolution":"\"5472000/899\"","FocalPlaneYResolution":"\"3648000/599\"","GPSInfoIFDPointer":"9894","GPSVersionID":"[2,3,0,0]","ISOSpeedRatings":"800","InteroperabilityIFDPointer":"9864","InteroperabilityIndex":"\"R98\"","Make":"\"Canon\"","MakerNote":"\"\"","MeteringMode":"5","Model":"\"Canon EOS 7D Mark II\"","Orientation":"1","PixelXDimension":"5472","PixelYDimension":"3648","ResolutionUnit":"2","SceneCaptureType":"0","ShutterSpeedValue":"\"679936/65536\"","SubSecTime":"\"17\"","SubSecTimeDigitized":"\"17\"","SubSecTimeOriginal":"\"17\"","ThumbJPEGInterchangeFormat":"10006","ThumbJPEGInterchangeFormatLength":"6276","UserComment":"\"\"","WhiteBalance":"0","XResolution":"\"72/1\"","YCbCrPositioning":"2","YResolution":"\"72/1\""}%
```


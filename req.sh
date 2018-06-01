curl -vXPOST 'https://blog.rajatjindal.com/xmlrpc.php' \
  -H"Host: blog.rajatjindal.com" \
  -H"Accept: */*" \
  -H"Accept-Encoding: gzip, deflate" \
  -H"Accept-Language: en-us" \
  -H"Content-Length: 713" \
  -H"Content-Type: text/xml" \
  -d '<?xml version="1.0"?><methodCall><methodName>metaWeblog.newPost</methodName><params><param><value><i4>0</i4></value></param><param><value><string>some-username</string></value></param><param><value><string>some-thing-here</string></value></param><param><value><struct><member><name>title</name><value><string>Testblog</string></value></member><member><name>post_type</name><value><string>post</string></value></member><member><name>description</name><value><string>I am onlinenow</string></value></member><member><name>categories</name><value><array><data></data></array></value></member><member><name>post_status</name><value><string>draft</string></value></member></struct></value></param></params></methodCall>'

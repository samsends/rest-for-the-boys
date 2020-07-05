curl -i -X  POST http://localhost:8080/tags \
  -H "Accept: application/json" -H "Content-Type: application/json" \
  -d '{ "tag": "rest", "msg": "when i am dead"}'

echo "\n\ngoto: http://localhost:8080/tags/<YOUR_TAG>"
#include <Arduino_JSON.h>

const int sensorCount = 4;
const int pins[] = {1, 2, 3, 4};

void setup() {
  Serial.begin(9600);
  while(!Serial);
}

void loop() {
  JSONVar jsonObj;

  for(int i=0; i<sensorCount; i++) {
    jsonObj[i]["value"] = 10 + i;
  }

  String str = JSON.stringify(jsonObj);
  Serial.print(str);
}

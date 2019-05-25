#include <Arduino_JSON.h>

/*
 * 
 * 
 */

const int sensorCount = 4;
const int pins[] = {1, 2, 3, 4};
int input;

void setup() {
  Serial.begin(19200);
  while(!Serial);
}

void writeSensors() {
  JSONVar jsonObj;

  for(int i=0; i<sensorCount; i++) {
    jsonObj[i]["pin"] = i;
    jsonObj[i]["value"] = 10 + i;
  }

  String str = JSON.stringify(jsonObj);
  Serial.println(str);
}

void flushInput() {
  
}

void loop() {
  if (Serial.available()) {
    Serial.flush();
    input = Serial.read();

    if (input >= 0) {
      // Flush input
      while(Serial.read() >= 0){};
      writeSensors();
    }
  }

  delay(10);
}

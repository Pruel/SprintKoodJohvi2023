#!/bin/bash
CASE_ID=$(head -179 streets/Buckingham_Place | tail -1|cut -d"#" -f2);
echo $CASE_ID;
cat interviews/interview-${CASE_ID};
echo $MAIN_SUSPECT;
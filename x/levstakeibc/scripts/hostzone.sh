#!/bin/bash

# insert host zone
staykingd tx levstakeibc register-host-zone connection-0 cosmos uatom ibc/0429A217F7AFD21E67CABA80049DD56BB0380B77E9C58C831366D6626D42F399 channel-0 1 --from val --keyring-backend test
# get specific host zone
staykingd q levstakeibc show-host-zone localstayking
# get all host zone
staykingd q levstakeibc list-host-zone
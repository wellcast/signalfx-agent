if [[ -z "$KAFKA_BROKER" ]]; then
    echo "ERROR: missing mandatory config: KAFKA_BROKER"
    exit 1
fi
sleep 30
while true;
 do
   if [ "$KAFKA_VERSION" = "2.0.0" ]; then
     ./opt/kafka_2.11-"$KAFKA_VERSION"/bin/kafka-console-consumer.sh --bootstrap-server "$KAFKA_BROKER" --topic sfx-employee --max-messages $((10 +RANDOM % 100))
   else
     ./opt/kafka_2.11-"$KAFKA_VERSION"/bin/kafka-console-consumer.sh --bootstrap-server "$KAFKA_BROKER" --new-consumer --topic sfx-employee --max-messages $((10 +RANDOM % 100))
   fi
   sleep 5
 done

FROM ubuntu:16.04
COPY sms-service /sms-service
RUN chmod +x /sms-service
CMD /sms-service
EXPOSE 8081
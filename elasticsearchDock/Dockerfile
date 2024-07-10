# Use the official Elasticsearch image as the base image
FROM docker.elastic.co/elasticsearch/elasticsearch:8.14.2

# Install the analysis plugins for Russian and Romanian
RUN elasticsearch-plugin install analysis-icu
RUN elasticsearch-plugin install analysis-phonetic
RUN elasticsearch-plugin install analysis-smartcn
RUN elasticsearch-plugin install analysis-stempel
RUN elasticsearch-plugin install analysis-kuromoji
RUN elasticsearch-plugin install analysis-nori

# Add any additional plugins you might need
# RUN elasticsearch-plugin install <additional-plugin>

# Set the necessary environment variables
ENV ELASTIC_PASSWORD=ELASTIC_PASSWORD
ENV discovery.type=single-node
ENV xpack.security.http.ssl.enabled=false
ENV xpack.license.self_generated.type=trial

# Expose the default Elasticsearch port
EXPOSE 9200
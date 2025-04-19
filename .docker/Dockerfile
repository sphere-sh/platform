FROM alpine:3.2 AS runner

# ----------------------------------------------------------------
#
# Installation and configuration of Sphere Ingestion Service.
#
# -----------------------------------------------------------------

# Create a group named `sphere-ingestion`
#
RUN addgroup -S sphere-ingestion

# Create user for Sphere Ingestion Service
#
RUN adduser -S -h /etc/sphere/ingestion -G sphere-ingestion sphere-ingestion



# ----------------------------------------------------------------
#
# Installation and configuration of Sphere Data Storage
#
# -----------------------------------------------------------------

# todo: Install / Configure Sphere Data Storage

# ----------------------------------------------------------------
#
# Installation and configuration of Sphere Management Service.
#
# -----------------------------------------------------------------

# todo: Install / Configure Sphere Management Service

# ----------------------------------------------------------------
#
# Installation and configuration of OpenRC.
#
# -----------------------------------------------------------------

# Install OpenRC as init system
#
RUN apk add --no-cache openrc




# Start OpenRC in foreground
#
CMD ["/sbin/init"]
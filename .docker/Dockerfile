FROM alpine:3.2 AS runner

ARG SPHERE_VERSION=0.1.0

# ----------------------------------------------------------------
#
# Download binaries from the Sphere Universe repository.
#
# -----------------------------------------------------------------

# Download from https://github.com/sphere-sh/platform/releases/download/v0.1.3/sphere-ingestion
RUN apk add --no-cache curl

RUN mkdir -p /var/sphere/universe
RUN mkdir -p /var/sphere/ingestion
RUN mkdir -p /var/sphere/management
RUN mkdir -p /var/sphere/storage

# Download Sphere Universe
#
RUN curl https://github.com/sphere-sh/platform/releases/download/v${SPHERE_VERSION}/sphere-universe -o /var/sphere/universe/sphere-universe

# Download Sphere Ingestion
#
RUN curl https://github.com/sphere-sh/platform/releases/download/v${SPHERE_VERSION}/sphere-ingestion -o /var/sphere/ingestion/sphere-ingestion

# Download Sphere Management
#
RUN curl https://github.com/sphere-sh/platform/releases/download/v${SPHERE_VERSION}/sphere-management -o /var/sphere/management/sphere-management

# Download Sphere Storage
#
RUN curl https://github.com/sphere-sh/platform/releases/download/v${SPHERE_VERSION}/sphere-storage -o /var/sphere/storage/sphere-storage


# ----------------------------------------------------------------
#
# Installation and configuration of OpenRC and dependencies.
#
# todo: reimplement installation and configuration of OpenRC, need to do some research on how to do it properly.
# -----------------------------------------------------------------

# Install OpenRC and other dependencies
#
RUN apk add --no-cache openrc iproute2

# Network configuration
#
RUN echo -e "auto lo\niface lo inet loopback" > /etc/network/interfaces # Configure loopback interface.
RUN mkdir -p /run /etc/network # Create run and network directories

# OpenRC configuration
#
RUN echo 'rc_logger="YES"' >> /etc/rc.conf                    # Enable logging for OpenRC
RUN echo 'rc_sys="container"' >> /etc/rc.conf                 # Set system type to container
RUN echo 'rc_provide="loopback net"' >> /etc/rc.conf          # Provide loopback and network
RUN echo 'rc_controller_cgroups="NO"' >> /etc/rc.conf         # Disable cgroups for controller
RUN echo 'rc_depend_strict="NO"' >> /etc/rc.conf              # Disable strict dependency checking
RUN echo 'rc_need="!net !dev !dev-mount !udev-mount !udev !udev-settle !udev-trigger"' >> /etc/rc.conf # Disable unnecessary dependencies

# Prevent OpenRC from starting a getty on tty
#
RUN sed -i -e 's/^tty/#&/' /etc/inittab

# ----------------------------------------------------------------
#
# Installation and configuration of Sphere Universe.
#
# -----------------------------------------------------------------

#RUN addgroup -S sphere-universe
#RUN adduser -S -h /etc/sphere/universe -G sphere-universe sphere-universe
#
## Copy binary from builder stage
##
#COPY --from=builder /build/universe/universe /bin/universe
#RUN chown -R sphere-universe:sphere-universe /bin/universe
#
## Setup OpenRC for Sphere Universe
##
#COPY .docker/openrc/init/sphere-universe.initd /etc/init.d/sphere-universe
#RUN chmod +x /etc/init.d/sphere-universe
#RUN rc-update add sphere-universe default
#
## Create log directory with proper permissions
#RUN mkdir -p /var/log && touch /var/log/sphere-universe.log && \
#    chown sphere-universe:sphere-universe /var/log/sphere-universe.log

# ----------------------------------------------------------------
#
# Installation and configuration of Sphere Ingestion Service.
#
# -----------------------------------------------------------------

#RUN addgroup -S sphere-ingestion
#RUN adduser -S -h /etc/sphere/ingestion -G sphere-ingestion sphere-ingestion

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

# Start OpenRC in foreground
#
CMD ["/sbin/init"]
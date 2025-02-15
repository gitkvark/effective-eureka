#/bin/bash

# Note: The test specified not to use LVM.
# The Rocky Linux 9 installer didn't offer the option, even though I was looking for it,
# as it was specified that it should be non-LVM.
# Full disk usage (apparently) means LVM for the installer...

# Grow the partition, if the underlying virtual disk size has changed.
/bin/growpart /dev/vda 2

# Extend the logical volume to occupy 100% of the available space.
/sbin/lvextend -l +100%FREE /dev/mapper/rl-root

# Grow the file system
/sbin/xfs_growfs /dev/mapper/rl-root

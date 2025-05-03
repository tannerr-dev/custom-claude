
i connected to my digital ocean droplet and set up ssh keys on my computer with the root user of my vps, after that i created a new user and logged out, now when i try to log in as the new user i get an error saying Permission denied (public key), what do i do now?
**1745618188520**

---
# Fixing the "Permission denied (public key)" Error for Your New User This is a common issue when setting up a new user on a VPS. Here's how to fix it:

## The Problem
You set up SSH keys for the root user but when you try to log in as your new user, authentication fails because:
1. The new user doesn't have your SSH public key in their authorized_keys file
2. SSH key authentication is required (password auth is likely disabled)

## Solution

You'll need to log back in as root and set up the SSH keys for your new user:

1. **Log back in as root**:
   ```
   ssh root@your_droplet_ip
   ```

2. **Set up the SSH directory for your new user**:
   ```
   mkdir -p /home/your_new_username/.ssh
   ```

3. **Copy your public key to the new user's authorized_keys file**:
   ```
   cp /root/.ssh/authorized_keys /home/your_new_username/.ssh/
   ```

4. **Fix permissions** (very important for SSH security):
   ```
   chown -R your_new_username:your_new_username /home/your_new_username/.ssh
   chmod 700 /home/your_new_username/.ssh
   chmod 600 /home/your_new_username/.ssh/authorized_keys
   ```

5. **Try logging in as your new user again**:
   ```
   ssh your_new_username@your_droplet_ip
   ```

If you can't log in as root anymore, you'll need to use the DigitalOcean console to access your droplet and fix the issue.
---
                     



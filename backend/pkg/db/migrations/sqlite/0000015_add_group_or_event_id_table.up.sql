-- 20240906_add_groupeoreventID_to_notifications.sql
ALTER TABLE Notifications
ADD COLUMN groupeId TEXT DEFAULT '';

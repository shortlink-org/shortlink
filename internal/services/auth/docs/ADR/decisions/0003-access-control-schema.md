# 3. Access Control Schema Definitions for Our Application using SpiceDB

Date: 2023-10-19

## Status

Accepted

## Context

This section describes the access control schema definitions for our application using SpiceDB.

> Playgrounds: https://play.authzed.com/s/uQrKpM5mUJ5o/expected

## Entities

+ **user:** Represents a registered user's account in our application.
+ **anonymoususer:** Represents an unauthenticated user.
+ **team:** Represents a team that can have members and administrators. Teams also have the ability to own links.
+ **link:** Represents a link with specified access control. Links can be owned by a team, read by various users, and written by specific users.

## Relationships and Permissions

**team:**
  + **administrator:** Indicates that a user is an admin of the team.
  + **member:** Indicates that a user is a member of the team.
  + **view_all_links:** A permission indicating whether a user can view all links in the team. Only administrators have this permission by default.

**link:**
  + **linkteam:** Indicates the team that owns the link.
  + **reader:** Indicates that the user (or team member, or any anonymous user) can read the link.
  + **writer:** Indicates that a user can write or modify the link.
  + **edit:** Permission to edit the link. Only writers have this permission.
  + **view:** Permission to view the link. Readers, writers, and any team with the `view_all_links` permission can view the link.

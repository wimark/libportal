# libportal changelog 

## [current]

 - Add: transferable flag to Authn config (381f0d3)
 - Add: enableLineCabinet bool (a19e7b0)
 - Add: SubscribeID and CreateAtSubscribe bson struct tag for (cbc8af2)
 - Add: SubscribeID and CreateAtSubscribe for account.go (9c76c3b)
 - Add: IsAuthConfirm and NextAuthConfirm for account.go (878d972)
 - Add: shortID field for portalProfile (8dd3893)
 - Add: subscription/locale types (c3cf510)

## [v1.6.0] 
 - WNE-2422: Add: add identity to PortalUserVoucher (88c4d9e)
 - WNE-2325: Add: PostAuthVSA in Access Server struct (4d4061c)
 - WNE-2325: Add: RDP intergration fields (229e306)
 - Del omitempty from cost (fd00317)
 - Add bson and omitempty for Cost PortalClientSession (70d0754)
 - WNE-2237: Add: const sms in PortalClientSession (61ac934)
 - WNE-2228: Add: "removal" of identification and authorization of clients (5a969e3)
 - WNE-2019: Add: Callfront as authN type (e0490a1)
 - WNE-2000: Add: user photo to account (f8e3fbf)

## [v1.5.0]

 - WNE-1975: Improve: golangci-lint issues (29750a0)
 - WNE-1963: Add: theme to page (d49a6ae)
 - WNE-1942: Fix: add PortalConditionRequest (4e01b2a)
 - WNE-1942: Fix: add loc_id bson tag (27eea3b)
 - WNE-1957: Add: PortalConditionRequest for condition (0acaee9)
 - WNE-1948: Add: Segment and Hash to structs (8868432)
 - WNE-1927: Add: PageURL to PortalPageSchedule (7a80469)
 - WNE-1919: Fix: package name (3021a12)
 - WNE-1919: Add: report enums flags (46af7f3)
 - WNE-1927: Add: Schedule/PageSchedule for LocID URL (30a6c28)
 - WNE-1892: Add: LocID to most of portal objects (d6e3861)
 - WNE-1892: Add: LocIDs to APIRequest (cd9e520)
 - WNE-1900: Add: Redirect Button Text and Redirect Button Color to ProfileAdData (938a8f0)
 - WNE-1900: Fix: delete tag json and bson in PollFields (0a3f9ca)
 - WNE-1900: Add: Redirect Button Text and Redirect Button Color to ProfileAdData (63c88a8)
 - WNE-1892: Change: rename GroupID to actual LocID (2af3d68)
 - WNE-1892: Add: GroupID as condition, Del oui/http_util (eda7fe1)
 - WNE-1869: Add: NasId to Ads log and Add Time (40990b8)
 - WNE-1869: Refactor: add ad follow url (bc525d9)
 - WNE-1872: Add: nas to portal client session (a171e85)
 - WNE-1869: Fix: id in ads log (07b18d6)
 - WNE-1869: Fix: Rename portal backend response and request (581e36f)
 - WNE-1869: Add: Account to ads log (d3163cb)
 - WNE-1869: Add: coll.go updates (50168d9)
 - WNE-1869: Rename: PortalBackendRequest (e3ed8ae)
 - WNE-1869: Add: PortalRequestObject (PortalBackendRequest) from libwimark (a9d1b82)
 - WNE-1869: Add: vendors file (695c528)
 - WNE-1869: Add: HTTP util function (8ea2e5b)
 - WNE-1869: Add: Redirect request obj (7a876bc)
 - WNE-1869: Add: Request to portal-backend (e3f614c)
 - WNE-1869: Add: Request (3d4c2f6)
 - WNE-1869: Add: collection (e91319c)
 - WNE-1869: Rename: AdLog to AdStatLog (c6b1b41)
 - WNE-1869: Add: Collection (5b85a6b)
 - WNE-1869: Add: User Data to AdLog (4f42894)
 - WNE-1869: Add: PortalAdLog (aabc101)
 - WNE-1863: Fix: Request type in APIRequest (ccff161)
 - WNE-1863: Del: excess fields in APIRequest (f7d4974)
 - WNE-1863: Add: APIResponse from portal-admin (47272a4)
 - WNE-1863: Add: APIRequest from portal-admin (588af21)
 - WNE-1864: Del: libtc (a517d1a)
 - WNE-1864: Add: fix pkg (b387f6a)
 - WNE-1864: Add: portal files (5a96c86)
 - Initial commit (bd42755)

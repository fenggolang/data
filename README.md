### Concepts for event sourced GNSS data handler

Could use DDB Table as event database

 -  Loose schema allows ..
    -  Primary key is StreamID - hashed uuid unique for some set of parameters (data source, data
       type, etc)
    -  Secondary key is EventNumber
    -  After that you just always have EventType (though I don't think DDB can enforce this)
 -  Process is ..
    -  Aquire write lock for a specific StreamID (lock entry is primary key = StreamID, secondary
       key = "lock" or something, value column is client ID or something?)
    -  Read last event for stream to get event number (can probably request top 1 item in
       descending order)
    -  Write next entry
    -  Release lock by deleting entry

A Site is an entity (unique receiver antenna pair) which has metadata and produces data. SiteMetadata
Aggregate may be more useful than any Data Aggregate. 

The Metadata may include a record of where data is, but streamed Observation Data itself is not
necessarily transactional - Observations shouldn't ever change (though it might be duplicated, in
the case of multiple streams from the same Site), so not sure if that fits into CQRS / ES.

Maybe there's a SiteData aggregate (or something) and it has the commands SubmitRinexFile,
SubmitRtcmObservation, SubmitSbfFile, SubmitSbfObservation, etc. then the original data is stored
in events and we project into the pure Observation structure.

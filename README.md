### Concepts for Normalized GNSS Observation Store

RTCM messages are not a particularly useful structure for doing anything with - but we can store
them in an append only database and project from there into a normalized read model.

Observations can't ever change - though might have multiple versions for the same observation, in
the case of multiple streams from the same Source (receiver and antenna pair), or receiving data
from the same Source in multiple formats.

700 sources producing on average 6 messages per second, averaging 300 bytes each in size = 4200
writes per second and around 1.3 MB per second.

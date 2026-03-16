# go-activities
Some take-home task-esque acitvities


# Task 1
### Public Facility Booking Service (Core Backend Logic)

This mimics things like community centre bookings, parking permits, or council facility reservations.

Task:<br/> 
Build a REST API in Go that manages bookings for public facilities.

Requirements: <br/>
Entities:

    Facility
        - id
        - name
        - location
<br>

    Booking
        - id
        - facility_id
        - start_time
        - end_time
        - booked_by
<br>

API endpoints:
    POST   /facilities <br/>
    GET    /facilities <br/>
    POST   /bookings <br/>
    GET    /bookings?facility_id= <br/>
    DELETE /bookings/{id} <br/>

Rules:<br/>
    Bookings cannot overlap for the same facility. <br/>
    Start time must be before end time. <br/>
    Return meaningful errors.

Example

    POST /bookings
        {
            "facility_id": "room1",
            "start_time": "2026-04-01T10:00",
            "end_time": "2026-04-01T11:00",
            "booked_by": "alice"
        }
# go-activities
Some take-home task-esque acitvities


# Task 1
### Public Facility Booking Service (Core Backend Logic)


### Task:<br/>
Build a small backend service in Go that allows users to create and manage bookings for shared public facilities (such as meeting rooms, sports halls, or community centres).

The system should expose a REST API where clients can:

register facilities

create bookings for a facility

list existing bookings

cancel bookings

The service must enforce business rules, particularly preventing overlapping bookings for the same facility and validating booking time ranges.

Requirements: <br/>
### Entities:

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

#### API endpoints:
    POST   /facilities <br/>
    GET    /facilities <br/>
    POST   /bookings <br/>
    GET    /bookings?facility_id= <br/>
    DELETE /bookings/{id} <br/>

### Rules:<br/>
    Bookings cannot overlap for the same facility. <br/>
    Start time must be before end time. <br/>
    Return meaningful errors.

### Example

    POST /bookings
        {
            "facility_id": "room1",
            "start_time": "2026-04-01T10:00",
            "end_time": "2026-04-01T11:00",
            "booked_by": "alice"
        }


# Task 2
## Case Management System

### Task: 

Implement a case management backend service in Go that tracks issues reported by users.

The system should allow clients to:

create new cases

retrieve case details

update the status of a case

attach comments to a case

Each case progresses through a defined lifecycle, and the service must enforce valid status transitions.

### Entities
```
Case
- id
- title
- description
- status (open, investigating, closed)
- created_at
```
<br/> 

```
Comment
- id
- case_id
- author
- message
```
<br/> 
  
#### API Endpoints
```
    POST   /cases
    GET    /cases
    GET    /cases/{id}
    POST   /cases/{id}/comments
    GET    /cases/{id}/comments
    PATCH  /cases/{id}/status
```
<br/> 
### Rules:

Status transitions must follow:

open -> investigating -> closed

Invalid transitions should return error.

### Example invalid:
```
open -> closed
```
Bonus task: <br/> 
Add filtering via:
```
GET /cases?status=open
```

# Task 3
### Data Processing Tool (Batch Processing)

### Task:
```
Build a CLI (Command Line)  tool that processes a CSV file of service requests.
You are encouraged to use Go Routines and Channels for this task

The tool should read the file, aggregate the data, and produce a summary report showing:
total cost per user
revenue per service category
```

### Example: 
```
CSV:

request_id,user,category,cost
1,Alice,parking,50
2,Bob,permit,20
3,Alice,parking,30
4,Charlie,permit,20
```
### Compute the following:


### Total cost per user

```
Example output:

Alice: 80
Bob: 20
Charlie: 20
```

### Top category by revenue

```
Example output:

parking: 80
permit: 40
Implementation constraints
```
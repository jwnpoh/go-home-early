*Go Home Early is currently in a hacky/as long as it works status. The app kinda works but is not guaranteed to work bug-free.*

# What
Go Home Early is a utility to help subject coordinators (especially those coordinating subjects with large cohorts and numerous classes/tutors) go home early.

# Why
Coordinators maintain a central document in Excel that keeps track of the cohort's results and data across all assessments in the two-year period. Each time there is an assessment there is a bunch of routine, repetitive actions that need to be taken. These are by no means difficult to do, but can be tedious and time-consuming.  

### Internal recording of marks
- individual tutors will take turns to update the central document in a local network folder
    - great risk of data corruption with many different users potentially messing things up
    - inefficient - bottleneck if one user forgets to close the file after he is done

**or**

- individual tutors fill up their respective marks according to a template, send to coordinator to cut-and-paste or vlookup to transfer to central document.
    - risk of error is reduced but not necessarily eliminated
    - labour-intensive, potentially time-consuming for the coordinator

### Uploading of marks to Cockpit
- individual tutors **must** upload themselves according to assigned classes
    - some tutors are averse to 'new technology', or don't trust their own ability to work with the preparation and uploading of csv files
    - tutors can key in directly in Cockpit but this is time-consuming and again, risk of human error
    - different marking pace. Coordinator might have to rush to generate and push out the csv files with marks back to tutors for uploading.

### We live in the 21st century
Why are we not automating some of these mundane routine tasks?


# How  

### Launch from terminal
```
$ go-home-early
```
### Launch from GUI
Double-click the executable file.

Both methods will launch an instance of the default browser requesting at `localhost:2021`.  
  
<br>


# Constraints:
- All files must be in csv format. If the central results list is maintained in an Excel file, export (save-as) the relevant sheet for the assessment to csv format. Possibly a future version that enables working directly with Excel (see to-do)
- Make sure that the names on the *central results list* match exactly as they appear on Cockpit. Order does not need to be the same. (similar to Excel vlookup).


# Progress:  
- [x] Basic server and routing
- [x] Generate command
- [x] Record command
- [ ] Cockpit command
- [ ] Comments
- [ ] Tests
- [ ] Analyse command
- [ ] Look into session management
- [ ] Look into security
- [ ] Better organisation of routes and services  


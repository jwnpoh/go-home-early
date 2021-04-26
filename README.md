# What
Go Home Early is a utility to help coordinators go home early.

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
```
$ gohome <command> [arguments]
```

The commands and their arguments are:
- `gohome generate 'central results sheet.csv'`  
        generates individual marks sheets for tutors by tutor name from 'central results sheet.csv'.  
- `gohome record 'central results sheet.csv' 'directory containing all tutors marks sheets'`  
        records marks from tutors' marks sheets onto central results sheet.  
- `gohome cockpit 'central results sheet.csv' 'directory containing all downloaded csv templates from Cockpit'`  
        populates csv templates provided by tutors with the total mark to be submitted in Cockpit.  
  
<br>


# Constraints:
- All files must be in csv format. If the central results list is maintained in an Excel file, export (save-as) the relevant sheet for the assessment to csv format. Possibly a future version that enables working directly with Excel (see to-do)
- Make sure that the names on the *central results list* match exactly as they appear on Cockpit. Order does not need to be the same. (similar to Excel vlookup).


# Limitations/to-do:  
- [x] Implement batch lookup from multiple input files to central results list.
- [x] Error handling. Should output message to inform user if there are records that did not match.
- [x] Add comments.
- [x] Get user input to indicate the correct column to match instead of hard-coding it.
- [x] Get user input to indicate which column to find the final mark value.
- [ ] Auto-detect the correct column to match instead of hard-coding it?
- [ ] Auto-detect which column to find the final mark value?
- [ ] Review the matching logic to be more rigorous. May need help.
- [ ] Tests. Still figuring out how tests work and how to write tests.
- [ ] Consider working directly with Excel files with the help of some Excel-Go packages.

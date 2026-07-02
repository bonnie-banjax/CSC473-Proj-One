
ANDREW MOBUS
LAB REPORT



# TITLE 
Overnight to Daytime Child Care Program Ratio in New York City
<!--|======[1]======[2]======[3]======|-->

# ABSTRACT

This paper seeks to assess the ratio of overnight child care service offerings
in New York City via cross reference of publically available government records.
Due to incomplete and disjointed information, much more time and effort was 
deployed towards attempts at data science than originally intende per the scope
of the report, compromising efficacy and completeness of results. Based on data
collected from OpenDataNYC (registered child care porviders dataset, PLUTO)
and information pulled from the the New York State Office of Child Care 
Services, there are roughly ~780 facilities opporated through 580 different 
organizations which offer non-traditional hours, which makes this represent
the maximum possible number of registered overnight child care providers out
of a total of just shy of 3000.



# INTRODUCTION

New York City is known as "the city that never sleeps", but someone has to watch
the children of those working the nightshift. Often, this falls on spouses or
informal networks of support; for a city of seven million non-commuting 
residents, there must be substantial formal programs available in order to
meet the needs of the city's off-hours labor. How one might deduce which of the
close to 3000 registered child care programs can watch one's children, and on
what basis of forward notice (ie, dealing with last minute shift changes) 
remains challenging, and requires non-trivial data aggregation efforts and
validation. Fortunately, the city and state of New York publish extensive public
data sets regarding all manner of areas of note, the relevant ones for this
study being the PLUTO (Primary Land Use and Tax Output) dataset, which was used
as an attempted means of validating capability opr categorical exclusion from
overnight oepration via Certificates of Occupancy and Building Class, but these
efforts resulted in relatively little to no useful insight, and then the 
registry operated by the New York State Office of Child Care Services, as well
as the New York City dataset of registered child care programs.


<!--|======[1]======[2]======[3]======|-->





# MATERIALS & METHODS
<!-- ORD -->
Identifying which of New York City's close to three thousand child care programs
offer overnight services necessitated a methodocal and data-driven approach; 
official government resources constitute almost the entity of the materials
consulted. Some data was readily accessible, as was the case for the current
registered child care programs; the first step was identifying which city 
government department had jurisdiction over such programs-- for NYC specifically
this falls on the New York City Department of Health and Mental Hygiene-- and 
searching the NYC Open Data repository. There is a data set covering register
child care programs, which provided the baseline total statistic. The next step
was identifying potential overlapping datasets which might provide some
means of establishing whether any of the providers were categorical excluded
from operating at night, based on building class designation or Certificate of 
Occupancy. Unfortunately, this proved to be a costly side quest, requiring to
pull in the PLUTO (Primary Land Use Tax-Lot Output) dataset, and attempt to 
join the datasets on the Borough-Block-Lot (BBL) field, as, remarkably, the
PLUTO dataset lacks the city's Building Identificatio Number (BIN) as one of
its over 100 columns. There was additional confusion as the "Building Class"
term used in city datasets has two different defintions based on whether it is
from the Department of Finance (DOF) or the Department of Buildings (DOB).
 
While joining the initial registered child care programs dataset with PLUTO
allow for precursory analysis of the various building types in which these
programs operate, the granularity was insufficient to provide anything close to
a robust or resonable guess as to how many might offer overnight care. 
Therefore, further data sources were sought out, this time from the New York
State Office of Child Care Services, which keeps a registry of all child care
providers in the state. The full dataset was acquired and then filtered on the
overseeing entity, which as established before, is NYCHMH; the results would
be cross-indexed with the city's provide registry at a later step. The New York
State OCFS has, on its web services page, a search tool, which has a query field
for "Non-Traditional Hours", which is defined on said web portal as follows:

```md
# Traditional Hours of Operation:

  - Day Care Centers, Small Day Care Centers, Family Day Care, and Group Day 
  Care traditional hours are Monday through Friday from 6:00 AM to 7:00 PM.
  - School Age Child Care Programs traditional hours are Monday through Friday 
  two hours before and three hours after school, when school is in session, and
  full days following regular opening and closing times during school breaks 
  and vacations.

# Non-Traditional Hours
Non-Traditional Hours of Operation:
Anything outside Traditional Hours of Operation is Non-Traditional Hours.
```

This statistic is recorded on the register program's index page, but, crucially,
not in the dataset offered by the OCFS. I initially explored setting up a basic
web scrapping bot in order to collect all the pages, as the URL schema was 
predictable and the identification number needed for such was in the dataset,
but ultimately, a more manual, low-tech solution was taken: use the search 
portal's filter on County/Borough and "Offers Non-Traditional Hours" toggle on
each of the five boroughs, seperately, blindly grabbing all text off those 
pages, and then extracting the predictable Licence/Registration number from all
those aggregated text blobs, which could then be turned into a table in the
database I was using, and joined against the OCFS dataset previously acquired.
This could then further be joined based on program name to the original NYC
dataset of registered child care programs.

While further data analysis might yield more accurate figures, due to time
constraints and the parameters of this project not necessarily being one
concerned with data science, the filtering has to be cut short admidst two
seperate botched attempts to perform the final validation join; additionally,
as might be noted from the cited definition of "Non-Traditional Hours", these
would not necessarily all be programs which offer overnight care, but rather,
represent the maximum possible programs which could provide such services;
there are some further options available for narrowing down the possible
maximum pool, suc has the difference between group care and school-based care,
as well as revisiting the building class and similar auxiliary data metrics
given the more focused candidate pool, the results stand where they are at a
maximum of roughly ~780 facilities opporated through 580 different 
organizations, ie, some programs have more than one site where they operate.
<!-- END -->

<!--|======[1]======[2]======[3]======[>X<]======|-->

# RESULTS
Roughly ~780 facilities opporate through 580 different organizations across
the five boroughs which are identified as offering non-traditional hours;
as already stated, this represents the maximum number of possible overnight
programs, but it does not guarantee that such full night hours are provided.
Ultimately, the analysis was inconclusive and requires substantially more
work to validate and cross reference the findings, as well as identify further
means of interrogating available data.

# DISCUSSION
My initial expectation that there were relatively few overnight offerings, as
had previously been suggested by a prior research project, appears to have been
mistaken, and is a strong reminder to make sure one fact-checks suprisingly 
claims, even if they come from one's teamamtes. While the city's information
services are, all considered, quite good, the reality is that the disjointed yet
overlapping data profiles ought to be better standardized across similar offices
rather than eveyrone making up their own spreadsheet schema; while the data is
usually provided as some form of CSV or spreadsheet, it is unknown whether these
are the actual storage formats as systems used by the primary databases or if 
they are mearly provided in this format for convience.

# CONCLUSION

Despite its moniker, New York City is far from as active at night as it could
be; the initial motivation for this studfy was as part of an ongoing assessment
of the efficacy and blocking problems on the city's Off Hours Delivery program,
which is successful but deals primarily in providing captial loans for the
purchase of automation equipment, rather than assisting in source labor for
unfilled roles which might command anywhere between a 5%-20% premium of the 
daytime baseline.



<!--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-->
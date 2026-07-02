
<style>
  :root {
    --bg-color: #0d0d0d;
    --surface-color: #1a1a1a;
    --text-main: #e0e0e0;
    --text-dim: #8f8f9a;
    
    --neon-teal: #00f3ff;
    --neon-pink: #ff007f;
    
    --electric-yellow: #ffee00;
    --thermal-orange: #ff3c00;
    --digital-grape: #7b2cbf;
    
    --terminal-green: #00ff41;
    --terminal-bg: #001100;
    --terminal-bdr:#00aa41;
    
    --og-pre-bg: #050505;
    
    --tble-bg: #000;
    --tble-bdr: #333;
    --th-bg: #222;
    --td-bg: #111;
    --td-txt: #fff;
    --tr-td-hvr-bg: #1a1a1a;
    
    --body-mark-bdr: #3e4446;
    --body-mark-grdnt-1: #13131f;
    --body-mark-grdnt-2: #050505;
  }

  /* Base Reset */
  html, body {
      background-color: var(--bg-color) !important;
      color: var(--text-main) !important;
      font-family: 'Segoe UI', Roboto, 'Helvetica Neue', sans-serif;
      line-height: 1.6;
      margin: 0;
      padding: 2rem;
      display: flex;
      flex-direction: column; /* This forces the vertical stack */
      align-items: center;    /* This centers the blocks horizontally */
  }
  }


  
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  /* Container for Pandoc content */
  body > * {
      max-width: 1000px;
      width: 100%;
  }

  /* 1. The "---" Breaks: Turning them into Digital Scans */
  hr {
      border: none;
      height: 2px;
      background: linear-gradient(90deg, transparent, var(--neon-pink), var(--neon-teal), transparent);
      margin: 3rem 0;
      position: relative;
  }
  hr::before {
      content: "SCANNING_SECTION_BOUNDARY...";
      position: absolute;
      top: -15px;
      left: 50%;
      transform: translateX(-50%);
      font-size: 0.6rem;
      color: var(--neon-pink);
      letter-spacing: 3px;
  }
  
/*  hr:after {
      content: "////////////////////////////////////////";
      display: block;
      color: var(--neon-pink);
      font-size: 8px;
      text-align: center;
      letter-spacing: 4px;
  }*/
  
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  
  
  /* 2. List Item "Marks" and Bold First Words */
  /* This targets the "The Assumption:" part of your lists */
  li strong:first-child {
      color: var(--neon-pink);
      text-transform: uppercase;
      font-family: monospace;
      background: rgba(255, 0, 127, 0.1);
      padding: 0 4px;
      border-right: 2px solid var(--digital-grape);
  }

  /* Custom Bullet Points */
  ul {
      list-style: none;
      padding-left: 1.5rem;
  }
  li::before {
      content: "»";
      color: var(--neon-teal);
      margin-right: 10px;
      font-weight: bold;
      text-shadow: 0 0 5px var(--neon-teal);
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  /* Typography */
  h1, h2 {
      color: var(--neon-teal);
      text-transform: uppercase;
      letter-spacing: 2px;
      text-shadow: 0 0 10px rgba(0, 243, 255, 0.5);
      border-bottom: 2px solid var(--neon-pink);
      padding-bottom: 0.3rem;
      margin-top: 2rem;
  }

  h2:hover {
      text-shadow: 
          2px 0 #ff007f, 
          -2px 0 #00f3ff;
      /*animation: glitch-anim 0.7s infinite;*/
  }

  @keyframes glitch-anim {
      0% { transform: translate(0); }
      20% { transform: translate(-2px, 2px); }
      40% { transform: translate(-2px, -2px); }
      60% { transform: translate(2px, 2px); }
      80% { transform: translate(2px, -2px); }
      100% { transform: translate(0); }
  }
  
    /* 3. The "##" vs "###" Hierarchy Debug */
  h3 {
      border-left: 10px solid var(--neon-teal);
      border-bottom: 1px solid var(--neon-pink);
      padding-left: 15px;
      background: linear-gradient(90deg, rgba(0, 243, 255, 0.1), transparent);
      text-shadow: 0 0 10px rgba(0, 243, 255, 0.5);
  }
  h3:hover {
      text-shadow: 
          2px 0 #ff007f, 
          -2px 0 #00f3ff;
/*       animation: glitch-anim 0.7s infinite; */
  }
  
  h4 {
      color: var(--electric-yellow) !important;
      font-family: monospace;
      text-transform: uppercase;
      background: rgba(255, 238, 0, 0.1);
      padding: 8px 15px;
      border-left: 4px solid var(--electric-yellow) ;
      width: fit-content;
      position: relative;
      margin-top: 2rem;
  }

  h4::after {
      content: "";
      font-size: 0.6rem;
      opacity: 0.7;
  }
  h4:hover {
      text-shadow: 
          2px 0 var(--thermal-orange), 
          -2px 0 #00f3ff;
      /*animation: glitch-anim 0.7s infinite;*/
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  a {
      color: var(--neon-pink) !important;
      text-decoration: none;
      transition: all 0.3s ease;
      font-weight: bold;
  }

  a:hover {
      color: var(--neon-teal) !important;
      text-shadow: 0 0 8px var(--neon-teal);
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  /* Code & Preformatted */
  code {
      background-color: var(--surface-color);
      color: var(--text-dim);
      padding: 0.2rem 0.4rem;
      border-radius: 3px;
      font-family: 'Cascadia Code', 'Fira Code', monospace;
  }

  pre {
      background-color: var(--og-pre-bg) !important;
      border: 1px solid var(--body-mark-bdr);
      box-shadow: 0 0 5px rgba(0, 243, 255, 0.1);
      backdrop-filter: blur(5px);
      padding: 1.5rem;
      overflow-x: auto;
      border-left: 5px solid var(--neon-teal);
  }

  pre code {
      background-color: transparent;
      padding: 0;
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/


  
  table {
      border-collapse: separate;
      border-spacing: 2px;
      background-color: var(--tble-bg);
      border: 1px solid var(--tble-bdr);
      width: 100%;
      font-family: 'Share Tech Mono', monospace; /* Very 'data' look */
  }

  th {
      background: var(--th-bg) !important;
      color: var(--electric-yellow) !important;
      text-transform: uppercase;
      letter-spacing: 2px;
      border-bottom: 2px solid var(--electric-yellow);
      padding: 12px;
      text-align: left;
  }

  td {
      background: var(--td-bg);
      border: 1px solid var(--th-bg);
      padding: 10px;
      color: var(--td-txt);
  }

  tr:hover td {
      background: var(--tr-td-hvr-bg);
      color: var(--neon-teal);
      border-color: var(--thermal-orange);
      cursor: crosshair;
  }

  /* Column Header Diagonal Cut */
  th:first-child {
      clip-path: polygon(0 0, 100% 0, 100% 100%, 15% 100%, 0 75%);
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  /* Blockquotes */
  blockquote {
      border-left: 4px solid var(--digital-grape);
      margin-left: 0;
      padding-left: 1.5rem;
      font-style: italic;
      color: var(--text-dim);
      background: rgba(0, 243, 255, 0.05);
      padding-top: 0.5rem;
      padding-bottom: 0.5rem;
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  /* Custom "Turn" Classes for your specific project */
  .body-mark {
      border-left: 4px solid #3e4446;
      width: 100%;
      border: 1px solid var(--tble-bdr);
      box-sizing: border-box;
      margin-bottom: 3rem;
      padding: 1rem;
      background: linear-gradient(145deg, var(--body-mark-grdnt-1), var(--body-mark-grdnt-2));
      box-shadow: 5px 5px 0px var(--neon-pink);
      position: relative;
      overflow: visible; /* Required for sticky children to work */
  }

  /* We use a real div or a specific header if you can, but as a CSS pseudo-element: */
  .body-mark::before {
      content: "LOG_ID: " attr(data-turn-id);
      position: -webkit-sticky;
      position: sticky;
      top: 10px;
      z-index: 10;
      
      /* Aesthetics */
      float: right;
      background: var(--digital-grape);
      color: #fff;
      font-family: monospace;
      font-weight: bold;
      font-size: 0.7rem;
      padding: 4px 12px;
      margin-right: -100px; /* Slight offset to look "attached" to the outside */
      clip-path: polygon(0% 0%, 100% 0%, 90% 100%, 10% 100%);
      box-shadow: 0 0 10px var(--neon-pink);
  }
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/

  
  .turn-segment[data-type="query"] {
      background-color: var(--terminal-bg);
      box-shadow: 0px 0px 5px var(--terminal-bdr);
      border: 2px solid var(--terminal-bdr);
      color: var(--terminal-green) !important;
      font-family: 'Courier New', Courier, monospace;
      padding: 1.5rem;
      position: relative;
      overflow: hidden;
      box-shadow: inset 0 0 15px rgba(0, 255, 65, 0.2), 0 0 10px rgba(0, 255, 65, 0.1);
  }

  /* The Scanline Effect */
  .turn-segment[data-type="query"]::before {
      content: " ";
      display: block;
      position: absolute;
      top: 0; left: 0; bottom: 0; right: 0;
      background: linear-gradient(
          rgba(18, 16, 16, 0) 50%, 
          rgba(0, 0, 0, 0.25) 50%
      ), linear-gradient(
          90deg, 
          rgba(255, 0, 0, 0.06), 
          rgba(0, 255, 0, 0.02), 
          rgba(0, 0, 255, 0.06)
      );
      background-size: 100% 2px, 3px 100%;
      pointer-events: none;
      z-index: 2;
  }

  /* The "Typewriter" blinking cursor at the end of the text */
  .turn-segment[data-type="query"]::after {
      content: "_";
      animation: blink 1s step-end infinite;
  }

  @keyframes blink {
      from, to { color: transparent; }
      50% { color: var(--terminal-green); }
  }

  
/*--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-*/
  /* Scrollbar - The finishing touch */
  * {
      scrollbar-width: thin;
      scrollbar-color: var(--electric-yellow)  var(--bg-color);
  }
</style>

<div class="body-mark" data-turn-id="5">

<div class="turn-segment" data-type="query">

ANDREW MOBUS
LAB REPORT



# TITLE 
Overnight to Daytime Child Care Program Ratio in New York City

</div>
<div class="turn-segment" data-type="response">
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


</div></div>
<!--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|-->
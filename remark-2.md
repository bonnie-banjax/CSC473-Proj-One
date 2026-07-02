
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
**TO:** Mayor’s Office of Child Care and Early Childhood Education  
**FROM:** Division of Early Childhood Education (NYCPS)   
**Date:** June 12, 2026  
**Subject:** Proposal for a Two-Year Pilot Program Providing Affordable Overnight Childcare in High-Need NYC Communities
</div>

<div class="turn-segment" data-type="response">
---

**Proposal**  
Many NYC parents work nontraditional hours, including overnight shifts, yet have limited access to reliable and affordable overnight childcare. Overnight babysitters can be expensive and difficult to secure on short notice. The lack of institutionally licensed overnight childcare creates financial strain for families, negatively affects child wellbeing, and may reduce workforce participation and job retention**.** Through the launch of a two-year pilot program partnering with existing licensed daycare facilities, affordable overnight childcare can be provided to communities in high need.

**Childcare shortages Limit Workforce Participation**  
According to a 2022 blueprint from administration of Mayor Eric Adams more than 500,000 New York City residents didn't seek employment in mid-2021 because of unmet child care needs (New York City Mayor's Office 9). Currently, there are not enough affordable seats for the city’s 500,000 children under five, and existing childcare options fail to meet the full set of childcare needs. These gaps force families to make difficult decisions that leave them underemployed, cause them to pass on career advancement opportunities, and prompt them to leave the workforce altogether (New York State and City of New York 121).  

**Existing Childcare Doesn't Meet The Needs of Non-traditional Hour Workers**  
The Shift Project’s Secure Scheduling study shows that 64% of workers nationally receive less than two weeks notice of their work schedule and 57% of workers experience at least one last-minute shift timing change per month. An estimated 43% of children in the United States have at least one parent that works non-traditional schedules; however, just 8% of centers and a third of home-based programs are open outside normal working hours. State child care regulations often don’t provide guidance on what quality care looks like early in the morning or late in the evening. Only 5% of registered daycare providers operate between 6 p.m. and 8 a.m., despite 780,000 parents in NYC working those hours. As a result, existing licensed seats fail to meet the demands of parents who work nonstandard hours (New York State and City of New York 121-122). 

**Childcare Shortages Harm Employers & Economy**  
The U.S. Chamber Foundation reveals that absenteeism and employee turnover cost firms anywhere from $400 million to $3 billion annually. Businesses may not be able to adequately staff operations, resulting in facility closures or restricted hours of operation across all industries (Barfield).

**Communities With The Highest Need**  
1 million New Yorkers make up the frontline workforce accounting for roughly 25% of the city’s total workforce with dense clusters living in specific neighborhoods across Brooklyn, Queens, and the Bronx. The New York City’s Frontline Workers report identifies 11 neighborhoods that each housed more than 20,000 frontline workers, the neighborhoods with the highest numbers include;

* Brooklyn: Canarsie, Flatlands, East Flatbush, and East New York  
* Queens: Jamaica, Hollis, Queens Village, and Cambria Heights  
* Bronx: Castle Hill and Parkchester  
* Manhattan: Washington Heights and Inwood


**PROPOSED PILOT PROGRAM:**

**Implementation Plan**  
This two-year pilot program is proposed by the Division of Early Childhood Education, operating using existing licensed daycares located near public transportation. One facility will run with extended hours in neighborhoods with the most need: Canarsie/Flatlands (Brooklyn), Jamaica/Hollis (Queens), Washington Heights/Inwood (Manhattan), and Castle Hill/Parkchester (Bronx).

**Using Existing Resources**  
Collaborating with operating licensed daycares can provide support to the 7,000+ childcare providers struggling to stay afloat (New York State and City of New York 122). These daycares would provide lower initial startup costs and expedite implementation of programming, since these facilities already meet state and city childcare regulations and operational experience. These facilities are also equipped with child-sized furniture, sleeping areas and equipment, age-appropriate toys and educational materials, kitchen or food preparation areas, bathrooms designed for young children, and security systems and controlled entrances, amongst other necessities. Participating facilities must also comply with enhanced overnight safety standards that will be continuously improved as the continuation of the pilot. 

**Program Operations**   
This pilot will aim to provide care for children 6 months old through elementary school. Enrollment eligibility will be for parents/guardians working non-traditional hours prioritized by income and families with the greatest childcare need. Operating hours are designed around common overnight work schedules and commuting needs. Facilities will have a maximum enrollment capacity determined by building occupancy limits, staff-to-child ratios, and state childcare licensing regulations. Opening hours are designed around common overnight work schedules and commuting needs. Facilities will maintain defined enrollment limits, allowing flexibility for emergency scheduling changes while ensuring providers can accurately anticipate attendance despite fluctuations. A centralized online platform and mobile app could be used to manage applications, enrollment, attendance, waitlists, and communication between families and providers. 

**Staffing Model**  
Given that this is not an educational role, there is a broad pool of potential staff from which to draw. Staffing can include certified teachers not currently employed full-time, retired educators, substitute teachers, and licensed childcare workers. Non-educational staff may include nurses, health aides, and security personnel. All staff will meet applicable state and city certifications, licensing and employment requirements. Staff must undergo background checks, complete program-specific training, and participate in ongoing professional development. Due to non-traditional hours, the initiative will offer competitive overnight pay and flexible scheduling arrangements to attract and retain qualified overnight staff. Staff currently working at participating facilities are not required to participate but may choose to do so, as long as all requirements are met and schedules are properly adjusted. 

**Funding & Cost**  
M\&M Daycare, a 24 hour daycare in the South Bronx, demonstrates both the demand for overnight childcare and the expenses required to maintain these services. Budget constraints are the primary obstacle preventing daycares which operate on thin margins, from offering extended hours. Given the recent approval by the Office of Children and Family Services (OCFS) of NYCPS request to increase the pay differential for providers from 5% to 15%, the city has demonstrated a willingness to invest in expanded childcare services. Additional funding or grant provided through the DoE, state, and city as incentives for providers to offer extended care can increase service for New York City families in need (Hashem). Existing childcare subsidies may be used where applicable, along with anticipated employer partnerships with hospitals and transit agencies, with an application form for employers to participate. A sliding-scale cost structure based on household income, similar to the MTA Fair Fares program for low-income residents, can be implemented to help fund the program while maintaining low childcare costs. 

The launch of the two-year overnight pilot program by the New York City Department of Education would address a critical gap in childcare services for families working nontraditional hours. Families will benefit from reliable, affordable overnight childcare, and greater job stability. Children will have safe, supervised environments, with consistent routines. Employers will have improved staff morale and retention and increased workforce participation. This allows the city to support essential workers and strengthen the local economy.

Submitted by:

Najlae Homman, Areeba Ali, Andrew Mobus

**References**  
New York City Office of the Mayor. (2022, June). *Accessible, affordable, and high-quality child care and early childhood education in New York City: A blueprint for child care & early childhood education in New York City*. [https://www.nyc.gov/assets/home/downloads/pdf/office-of-the-mayor/2022/Childcare-Plan.pdf](https://www.nyc.gov/assets/home/downloads/pdf/office-of-the-mayor/2022/Childcare-Plan.pdf)

New York State Office of the Governor, & New York City Office of the Mayor. (2022, December). *New NY action plan: Making New York Work for Everyone*. [https://www.nyc.gov/assets/planning/download/pdf/plans-studies/office-reuse-task-force/New-NY-Action-Plan-Making\_New\_York\_Work\_for\_Everyone.pdf](https://www.nyc.gov/assets/planning/download/pdf/plans-studies/office-reuse-task-force/New-NY-Action-Plan-Making_New_York_Work_for_Everyone.pdf)

*Secure scheduling \- the shift project*. (2025, June 5). The Shift Project. [https://shift.hks.harvard.edu/secure-scheduling/](https://shift.hks.harvard.edu/secure-scheduling/)

Barfield, J. (2023b, October 26). *Billion dollar problem: working parents leave workforce, postpone school due to childcare challenges*. U.S. Chamber of Commerce Foundation. [https://www.uschamberfoundation.org/education/billion-dollar-problem-working-parents-leave-workforce-postpone-school-due-childcare](https://www.uschamberfoundation.org/education/billion-dollar-problem-working-parents-leave-workforce-postpone-school-due-childcare)

Office of the New York City Comptroller. (2020, March 26). *New York City's frontline workers*. [https://comptroller.nyc.gov/reports/new-york-citys-frontline-workers/](https://comptroller.nyc.gov/reports/new-york-citys-frontline-workers/)

Hashem, K. (2024, April 17). *Extended-Hours child care gives economic boost to South Bronx neighborhood*. Columbia News Service. [https://columbianewsservice.com/2023/12/11/extended-hours-child-care-gives-economic-boost-to-south-bronx-neighborhood/](https://columbianewsservice.com/2023/12/11/extended-hours-child-care-gives-economic-boost-to-south-bronx-neighborhood/)

Harknett, K., Schneider, D., & Luhr, S. (2022). Who Cares if Parents have Unpredictable Work Schedules?: Just-in-Time Work Schedules and Child Care Arrangements. *Social Problems*, *69*(1), 164–183. [https://www.jstor.org/stable/27393933](https://www.jstor.org/stable/27393933)

**AI Disclaimer** 

When writing this memo, the QuillBot rephraser tool was used. However, those portions were later re-paraphrased. An online tool GrammarChecks free check were used to improve grammar, spelling, and punctuation. Gemini was used for the citations of the PDF from the New York City Office of the Mayor and New York State Governor & New York state office PDF. Scribbr was used for the rest of the citations. 

Gemini was additionally used to locate sources (Shift Project and associated paper), as well as for fact checking and other auditing.

	 	

	 	

**Memo Format**

The memo was structured based on general principles while drafting and validated against the guidelines provided in the course textbook on pages 372-374.

**Contributions**

Najlae Homman: Main memo drafting & research for existing licences child care offerings in NYC.

Areeba Ali: Alternate memo drafting & original annotated bibliography formatting, final editing.

Andrew Mobus: Topic selection & research for existing care networks in underserved populations, editing.

The original mutual aid proposal was abandoned due to liability concerns, despite the centralization issues and cost overhead of the current proposal; ultimately, government must be transparent and accountable, and decentralized coordination services thrive by circumventing costly compliance requirements, making this approach fundamentally unsuitable where auditability is paramount, even if it were accepted on cost efficacy grounds.

</div></div>
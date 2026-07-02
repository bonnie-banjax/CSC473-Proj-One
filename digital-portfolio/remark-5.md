
<!--|======[1]======[2]======[3]======[>X<]======[5]======[6]======[7]======|--> 


<style>
/*.body-mark {
    border-left: 4px solid #3e4446;
    margin-bottom: 2rem;
    padding-left: 1rem;
}*/

/* Style all queries */
/*div[data-type="query"] {
    font-weight: bold;
    color: #aec2d3;
    background-color: #1b1e1f;
}*/

/* Style all responses */
/*div[data-type="response"] {
    font-family: serif;
    line-height: 1.6;
}*/


</style>

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
# SELF EVALUATION ESSAY

Too little, too late.  
Too all over the place.

</div>

<div class="turn-segment" data-type="response">
I don't think I'm making it in under the bell for submitting what needs to be submitted by 6pm tonight. I failed to keep pace with the course, and failed to maintain focus on the specific depth and scope required to meet such rapid deadlines. 

I'm proud of my engineering proposal, and were I just one day more time, I could have an outstanding end product-- but I'd still be down the entire lab report, which is arguably the most important assignment with regards to the procedural compliance and protocol alignment. Despite the emphasis on visual artifacts, I failed to produce any of note, or include them as any other than an afterthought or box to be checked. 

My explicit goal for this class was to address my weakness specifically in these two regards:
* The timeliness of submissions
* Adherence to the rubric to the letter

Those aren't the course learning outcomes, though. 

---

### Learning Outcome 6: Research Sources & Analysis

Despite having the least writing to show for the lab report out of all the assignments, the research phase aligns entirely with learning outcome 6, *"Locate research sources(...)"*; I was able to locate, after a very manual and intensive collection phase, all the materials I'd need to properly analyze and process the data in a manner which answered the core research question. I just couldn't turn it into a paper in time. 

I treated the assignment as an exercise in through and methodologically rigorous analysis, rather than one in producing an artifact which fit a highly structured and specific mold; the content was the least important thing about the report, in fact.

Was my failure one of being too logical? I don't think so. My failures were idealistic; a desire to use every assignment as an opportunity to push the bounds of what I am capable of, rather than shoring up the weaknesses of things which I can do, but do so clumisly, and poorly, and taking lots of time. 

> It took me three weeks to pivot the technical assignment's premise from one rooted in socio-political capture of society via unauditable technological interfaces to a piece undoubtably inspired by "Exhalation" exploring the loss of meaning as seen through the process of flattening two documentsm and the point at which they cease to be digital and become a corpse. 

Does that constitute examining how attitudes towards linguistic standards empower and oppress language users? It was explicitly inspired by identifying how the textual interface of programming languages defines the relationship of most people to computers, where either one speaks the magic language, and often mistakes the interface for the thing in of itself, or, for most people, find the rigid inflexibilies of programming languages something which bars access for all but they they might consider the magi; was that evident in the paper? Maybe.

---

### Learning Outcome 5: Audience & Visual Integration

Does my writing proficiency and extensive technical knowledge translate into addressing a wide range of audiences, by, say, recognizing what is essential technical information, and what is unnecessary as pertains to the larger point? 

**Not a chance.** I'm a one trick pony reliant on rheatorical flourish because I don't know-- or don't care to-- meet the reader where they are in terms of their technical knowledge and engagement. Not everyone wants to increase their level of technical literacy, and the purpose of technical writing is not necessarily to push technical literacy forward or foster engagement; often times, it is the opposite, and apply the right level of granularity such that the reader can engage with only that which is necessarily in that specific moment in order to understand what choice they want to make. 

My writing is so impenetrable that I don't even know if the technical components are as I think because no one who's read any of it has felt like they understood what I was trying to say. My failure to integrate diagrams natively in my work is likely my most damning failure for learning outcome 5 in terms of indicating the issue is a lack of will, not just a lack of ability.

---

### Collaboration, Citations, and the LLM Dilemma

I haven't written a single citation this entire semester. I've systematically tracked all of my sources, and have overflowing tabs of goverment data and academic reports, along with the LLM context windows from where I found links to such sources. I still vetted and interegated the sources and claims found through such tools, but I also failed to interrogate one of the basic claims established by a peer for the memo project-- that they only found one overnight child care service in new york city-- until substantially into the research phase. 

The collaboration for the memo project was mixed:
  1. We fulfilled our obligations piecemeal.
  2. Had minimal and passive aggresive group communications.
  3. The final memo is verbose, likely over-edited by an LLM writing assistance service.

Despite making a small fuss to my group about making sure our project had an LLM-usage disclaimer, I failed to provide one for the technical description; I used LLM's heavily in research, even if I don't use them for the final writing of assignments. I do fact check them, though, at least more than I typically did before the "slopocalypse"-- does that qualify for learning outcome 6? Even if it does, that's one of the outcomes already addressed before.

If I heavily use LLM's for research and intial material aggregation, but not for final wording synthesis, it probably means I'm implicitly engaging in summary, sythesis, and argumentation-- I 'spar' with LLM's as a means of rheatorically organizing my thoughts, but no matter how many times I end a query with *"Do not glaze."* I'm still engaging with a reflecting device pre-disposed to please, and most humans want to be told they're right. 

What have I learned, then?

---

### Reflection on Technical Elitism & Communication

Despite my egalitarian ideals, I'm elitist when it comes to technical affinity. I assume the only valid reason for people to be uninterested in increasing their technical proficiency is they are intimidated and feel that technology and the hard sciences are "not for them" because of poor pedagogy and general failure of the educational system. I assume that the problem is merely one of finding the right words to make the material engaging an entertaining; that if I can make the material accessible, people will be empowered to engage with it on their own terms. 

Not everyone gains the same marginal benefit and return on their time and effort when it comes to their particular life objectives and personal values. Much like my assignments failed because I over-engineered the standards to which I held them, my communication fails because I think simplifying things means I'm insulting my audience's intelligence, and what's almost always most important to me is trying to convince people to do better because that's what I want to do. 

---

### Classroom Dynamics & Final Compromises

At the start of the course, the professor (you, the reader) pointed out that, in the free write discussions the previous semester, there were often the same voices speaking with something to say, where perhaps the discussions would be richers were commentary more evenly distributed. There was never a discussion where I didn't have something to say nor didn't feel like I either was holding back or should have held back more. I still spoke too much, or maybe I just said the wrong things. 

Perhaps I'm not giving myself enough credit. 
  * My engineering proposal is directly inspired by listening to one of my peers articulate how much the current state of primary education troubled them. 
  * Even if I wasn't able to take my own advance about scope management, I was able to help at least one classmate focus down their lab report to something more manageable. 
  * When my group mates decided they wanted the proposal to pursue a centralized model of which everything I understood of the circumstance suggested would, were it implemented, be deeply flawed and contribute to the negative image of socialization resulting in money-pit social programs, I held my tongue and compromised. 

Still, when offered the chance to give the final words and reflection on the class, where everyone turned to me with the assumption that I would be the one to speak, all I said was:

> "no, I've yapped enough."

</div></div>
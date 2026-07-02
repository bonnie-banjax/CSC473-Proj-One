
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
  
  
  .pdf-container-1 {
    width: 100%;
    height: 20vh; /* Takes up 85% of the screen height */
    border: 1px solid #ccc;
    margin-bottom: 20px;
  } 
  
  .pdf-container-2 {
    width: 100%;
    height: 40vh; /* Takes up 85% of the screen height */
    border: 1px solid #ccc;
    margin-bottom: 20px;
  } 
  
</style>

<div class="body-mark" data-turn-id="5">

<div class="turn-segment" data-type="query">

ANDREW MOBUS
ENGL 21007 W4E
12026-06-15

# DOCUMENT TO DOCUMENT TO IMAGE TO PRINT: DOCX AND PDF PRINTING PROCESS

</div>

<div class="turn-segment" data-type="response">
What's the difference between a docx and a PDF digital document format? The former is a browser webpage masquerading as sheets of paper, while the latter is a parameterized mathematical model of what a printed document is. Yet, despite these at-rest differences, both are capable of being printed by most or practically all printers one might expect to encounter-- although for the former, this is only true by first turning it into a PDF. 

Their internal representations necessarily converge at the point of translation from declarative data structure to set of imperative sequential commands or true bit-by-bit decomposition such that every '1' maps to a single microscopic dot of ink precisely sprayed on a page. 

> Consider the creation of a clay vessel; at first, it is wet, workable, and unstable; when it dries, it is hard but not impervious to change, can readily be returned to workable state, and is ready for glaze and the kiln; once fired, the pottery is what it is, and all modification is purely destructive. 

We will map the digital document printing process and the differences between the aforementioned formats to this sequence of transformations and further map those to the processes by which computers operate, in broad scopes. The fancy name for this might be "reification", as in "the reification from program to process", but we'll stick with printing.

---

## The "Wet Clay": DOCX and the Hierarchical Tree

Our wet clay is the docx format, which, as a data structure, is a **hierarchical tree**; while the term might be unfamiliar, if you've used a file manager, you're intimately familiar with the key properties of trees. Think of every folder as being branches off of your main partition (a portion of your hard drive), and everything that isn't a folder is a leaf sprouting from those branches. More broadly, when one puts a folder or file inside of another folder, one is creating a parent-child relationship-- the child "extends from" or is contained in the parent, which might have its own parent, all the way back to the root. A thumb-drive/flash-drive is very much a potted plant version of this. 

Web pages are fundamentally built from a similar methodology of parent-child relationships; consider this paragraph as a parent to all the word elements inside of it, and each word is a parent to all the characters that make it up, while the characters are the leaf nodes of the whole system.

For a web browser, a paragraph is represented by `<p>` (and an ending tag, `</p>`), and for a word document, a paragraph is `<w:p>` (and an ending tag, `</w:p>`); while clearly similar, the docx file adheres to the Office Open XML standard, which is different from HTML, but they share the same origin and same abstract representation: the **Document Object Model**. Tags (the things in between angle brackets) are the "objects", and are a means of articulating the "structure" of a document independent of how it might layout on any given page, hence the reactive reflowing and scaling one commonly observes on webpages. 

Here is where our "wet clay" reveals itself: the docx file has no idea what shape it is, or what the concept of "shape" even means; it has no idea of its own formatting except insofar as the same tag means the same thing whenever it appears, and the sequencing of content. The entire document can be written on a single line and it will stil preserve its internal structure, but it relies entirely on its context-- the hands which shape it-- to provide it form and specifics.
</div></div>

<iframe src="tree-example.pdf" class="pdf-container-1"></iframe>

<!-- ![tree_diagram](./tree-example.pdf) -->
<div class="body-mark" data-turn-id="5">
<div class="turn-segment" data-type="response">
The tree component of docx files is not only in the "page level" Document Object Model (DOM) representation, but also in the storage of its resources; a docx file is just a specific kind of ZIP file, and one can crack one open with the same tools, revealing the potted plant inside. Even though the docx file does carry with it some resources and organizational structure (in the form of the file-system style parent-child folder-file relationships) along side the DOM representation, it is reliant on external forces to interpret the structure it does have and decide what, exactly, everything means; every time you print a docx file, the computer is playing a lego-bricks madlibs tetris game with your document and figuring out how to articulate what your document looks like at that particular point in time, with those particular settings, for this particular machine (the computer)-- all before anything remotely resembling what the printer will understand.

---

## The "Dried Clay": PDF and the Object Graph

Once the document has been anchored to a specific context, with whatever given font, margins, text kerning, layout, and a host of other characteristics, the docx file has, either in effect or literally, been turned into a PDF. For Windows, the transformation is less literal, while on Mac & Linux, it must, non-figuratively, be transformed into a literal PDF-- wet clay necessarily must dry for firing to take effect, even if it is being dried inside the kiln. 

So, what has changed in this transition? 
* For one, the document is no longer a hierarchical tree, but instead an **object graph**; 
* for another, it is now **spatially aware** and defined, using math, in explicitly spatial terms. 

Where the docx file was a ZIP archive-- which, it should be noted, has a "dictionary" of its contents located at the end of the file, as required by the standard-- the PDF is a stream of raw bytes, some of which are plain text (human readable), while others are binary.
</div></div>

<!-- ![graph_diagram](./graph-diagram.pdf) -->
<iframe src="graph-diagram.pdf" class="pdf-container-2"></iframe>

<div class="body-mark" data-turn-id="5">
<div class="turn-segment" data-type="response">
The resources which make up the PDF are woven together in a flat net or tangle of connections and references to each other; think about a Wikipedia article or classic webpage consisting of the local content of the page, with hyperlinks to other related articles scattered throughout. Those links turn the site into a graph, because a page isn't necessarily the parent of another page it has a link to, and the page that's linked to can have its own link right back causing you to loop in circles-- cycles-- if you aren't careful.

In the case of the PDF, this isn't exactly how the "graphness" of the format manifests; the file is "flat" due to being one long stream of bytes, and, much like the ZIP, there is a "dictionary"-- a collection of terms and what those terms refer to-- called the **Cross-Reference Table (XREF)** at the end, which tells whatever is reading the PDF where resources are in relation to one another in terms of byte offsets. While abstract, these byte offsets are a spatial representation of the relationships-- or edges, in math terms-- of the objects which make up a PDF. The types of those things-- the defined sorts or groupings which share behavior and "shape"-- are as followings:

> Booleans; Numbers; Strings; Name Objects; Arrays; Dictionaries; and Streams.

---

## Vector Graphics and Parameterization

The "data types" are also the core building blocks of PostScript, which is a **Page Description Language**-- a means of encoding and representing the physical, spatial relationships of a page with vector graphics, which allow the PDL description to apply to any given printing circumstance. It's worth noting this is not the same as with a docx format, where the fundamental spatial information is entirely missing (or rather, unspecified); instead, vector graphics mean the description is parameterized, meaning once you provide a specific value for a given term or set of terms, the rest of the terms will take on specific values based on the defined relationship between them. 

Scaling (zooming) is an example of this kind of coupled relationship; your screen has a fixed number of pixels, and when you zoom in or out, you are saying "Render that same thing, but now a basic unit of space is 1.25x times larger than it was a moment ago", which means that even though every single pixel might have a difference value than it did before, the image you see stays the same.

A PDF knows its page width and length, in an abstract sense, and knows the height and width of any given character in proportion; more than that, though, it knows a set of points and relationships of those points in such a way which say, for this given character, which occupies this given unit of space, this path (an equation describing a line or curve) should be painted, and therefore, right now, these specific pixels should have their values set to such and such a value. This applies to rendering the document to a given screen, but equally, the same logic applies to a page, which is just a screen with a very, very, very slow refresh rate. 

At each point, a number represents the particular output characteristics of that point-unit of space; the most basic version, that point only has (or needs) two values: "something here" (represented as 1), or "nothing here" (represented as zero). Whether that "thing" is light or ink doesn't really matter until one gets to the different color models (RGB, which is additive, versus CYMK, which is subtractive); either way, the structure of the 2D grid that constitutes the image remains the same.

```text
                            _________________
                            |X|O|X|O|O|X|X|X|
                            |X|O|X|X|O|X|O|X|
                            |X|O|X|O|O|X|X|X|
                            |X|O|X|O|O|X|X|O|
                            |X|O|O|O|O|X|X|X|
                            |X|O|X|O|O|X|X|X|
                            |X|O|X|O|X|X|X|X|
                            |X|O|X|O|O|X|X|X|

```

_Figure 1: a bitmap_

## Rasterization: The Final Collapse

Colloquially, the ASCII "diagram" of a 2D array is refered to as a "bitmap", and the resolution of that bitmap corresponds to the **Dots Per Inch (DPI)** of the printer; for higher end printers, the printer itself will construct this bitmap, through a series of imperative commands (which are a programming language known as PostScript, which uses the same data representation types as discussed with PDF's), and this is handled by the printer's drivers. For less sophisticated printers, the host computer must construct the bitmap and dump it in the printer's frame buffer.

The process by which the frame buffer is constructed is **"rasterization"**, which was alluded to earlier; given a set of points and the path relationships between them, assign values to unit partitions of space based on a given partitioning scheme.

Now, there is nothing left of the original document as one might understand it to mean-- there are no words, there are no fonts, no headings or footnotes or diagrams. Whether or not the physical printing has begun or finished, whether originally a docx file or a PDF, there is no distinction, because the artifact is not a digital document. Everything has been collapsed into a bespoke, static image for use with a specific analog screen.

</div></div>
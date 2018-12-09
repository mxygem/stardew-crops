Feature: Formatting options

  Available options:
    * Pretty - (default) When outputting data, this will display things in a table or in other nicely formatted ways.
    * JSON-Pretty - output formatted JSON
    * Raw - outputs non-formatted JSON

  Scenario: The pretty formatter returns nicely formatted data
    Given the pretty formatter is desired
    When info for garlic is requested
    Then the data must match
    '''
╔═══════════════╦═══════════════════════════════╗
║   + .+;++++   ║ Name: Garlic                  ║
║   ++;+.+. ++  ║ Seeds: Garlic Seeds           ║
║     +; ;;+    ║ Growth Time: 4 days           ║
║    +.;.;.+    ║ Season: Spring                ║
║   +.;. .;;+   ║ Continual: False              ║
║  +; ;.  ;.;+  ║ Seed Price:                   ║
║  +; ; . ; .+  ║   * General Store: 40g        ║
║  +; ;. .;.;+  ║   * JojaMart: N/A             ║
║   +.;...;;+   ║   * Traveling Cart: 100-1000g ║
║    +.;.;;+    ║   * Night Market: 40g         ║
║     +++++     ║ Note: Only avail Year 2+      ║    
╚═══════════════╩═══════════════════════════════╝
    '''

  Scenario: The json-pretty formatter returns json data with whitespace

  Scenario: The raw formatter returns json data with no whitespace


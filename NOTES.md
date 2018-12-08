# stardew-crops

default
    returns all crop data
    option(s)

        - format - string - default pretty
            - pretty - tabular data. ascii?
            - json-pretty - formatted json
            - raw - unformatted json
    info
        - look up specific crops by name
        - option(s):
            - format
    search
        search by:
            - crop name
            - season
                - csl of one or more seasons.
                - options: spring, summer, fall
            - harvest type
                - options: single, multi
            - growth time
                - specify time in days
        option(s):
            - format
            - expanded - bool - default true - shows all info data of matching crops

Questions:

- How are large data sets returned?

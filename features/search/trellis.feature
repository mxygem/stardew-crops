Feature: Search by Trellis

    Scenario: A list of crops grown on a trellis is desired
        # stardew-crops search --trellis true
        When a search for crops that grow on a trellis is performed
        Then a list of crops that grow on a trellis are returned

    Scenario: A list of crops not grown on a trellis is desired
        # stardew-crops search --trellis false
        When a search for crops that do not grow on a trellis is performed
        Then a list of crops that do not grow on a trellis are returned

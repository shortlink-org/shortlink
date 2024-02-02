<?php

use Behat\Behat\Tester\Exception\PendingException;
use Behat\Behat\Context\Context;
use Behat\Gherkin\Node\PyStringNode;
use Behat\Gherkin\Node\TableNode;

/**
 * Defines application features from the specific context.
 */
class FeatureContext implements Context
{
    /**
     * Initializes context.
     *
     * Every scenario gets its own context instance.
     * You can also pass arbitrary arguments to the
     * context constructor through behat.yml.
     */
    public function __construct()
    {
    }

    /**
     * @Given I am authenticated with the API
     */
    public function iAmAuthenticatedWithTheApi()
    {
        throw new PendingException();
    }

    /**
     * @When I send a POST request to :arg1 with the following data:
     */
    public function iSendAPostRequestToWithTheFollowingData($arg1, TableNode $table)
    {
        throw new PendingException();
    }

    /**
     * @Then the response code should be :arg1
     */
    public function theResponseCodeShouldBe($arg1)
    {
        throw new PendingException();
    }

    /**
     * @Then the response should contain the following JSON:
     */
    public function theResponseShouldContainTheFollowingJson(TableNode $table)
    {
        throw new PendingException();
    }

    /**
     * @When I send a GET request to :arg1
     */
    public function iSendAGetRequestTo($arg1)
    {
        throw new PendingException();
    }

    /**
     * @Then the response should contain a list of previous support requests in the following format:
     */
    public function theResponseShouldContainAListOfPreviousSupportRequestsInTheFollowingFormat(TableNode $table)
    {
        throw new PendingException();
    }
}

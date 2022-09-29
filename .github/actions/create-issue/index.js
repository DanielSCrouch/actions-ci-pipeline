'use strict';

const core = require('@actions/core');
const github = require('@actions/github');
const errorHandler = require('./src/error-handler');

const run = async () => {
  const token = core.getInput('token');
  const title = core.getInput('title');
  const body = core.getInput('body');
  const assignees = core.getInput('assignees');

  core.info(`token: ${token}`);
  core.info(`title: ${title}`);
  core.info(`body: ${body}`);
  core.info(`assignees: ${assignees}`);

  const octokit = github.getOctokit(token);
  const response = await octokit.rest.issues.create({
    ...github.context.repo,
    title: title,
    body: body,
    assignees: assignees ? assignees.split('\n') : undefined,
  });

  core.setOutput('issue', JSON.stringify(response.data));
};

run().catch(errorHandler);

#!/usr/bin/env node
var SwaggerToMarkdown = require('./swagger'), Parse = require('./args'), opt = Parse.opt;
var options = new Parse([
    opt('n', 'apiname', "Provide the API name", 'string'),
    opt('r', 'resourcefile', "Provide the resources.json to define your API resources filename", 'json_file', true),
    opt('p', 'parametersfile', "Provide the parameters.json to define your API parameters filename", 'json_file', false),
    opt('o', 'markdownfile', "Provide the api.md to define your output Markdown filename", 'filename'),
    opt('s', 'specifications', "List of specification files in a json format", 'json_file'),
    opt('h', 'help', 'Show this message', 'help'),
    opt('v', 'version', 'Show version')
]).parse(process.argv.slice(2));

console.log(new SwaggerToMarkdown().$enhance(options).print());

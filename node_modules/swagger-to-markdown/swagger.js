
function $titlelize(str) {
    if (!str) return;
    str.split(/[^a-zA-Z0-9]/).filter(function(v){
        return v;
    }).map(function (v) {
        return v.substring(0, 1).toUpperCase() + v.substring(1);
    }).join(' ');
}

function SwaggerToMarkdown() {
    var self = this;
    var _lines = [];
    this.print = function(){
        return _lines.join('');
    }
    var f = {
        $write:function () {
            for (var i = 0, l = arguments.length; i < l; i++) {
                _lines.push(arguments[i]);
            }
        }
    }

    this.$enhance = function (options) {
        var resources = options.resourcefile;
        this.parameters = options.parametersfile;
        this.$write_api_to_markdown(options.markdownfile, options.apiname, resources["apiVersion"], resources["basePath"], resources["apis"], options.specifications);
        return this;
    };

    this.$write_api_to_markdown = function (markdown_file, api_name, api_version, base_path, apis, specifications) {

        self.$write_header(f, api_name, api_version, base_path);
        return apis && apis.forEach(function (resource, index) {
            var path = resource.path + $titlelize(resource.resource);
            f.$write(self.$build_markdown_header(
                $titlelize(self.$extract_resource_name(resource.path)+'resource'), 2
            ))
            f.$write(self.$build_markdown_header(self.$extract_resource_name(path), 2));
            f.$write(resource.description+"\n\n");
            (Array.isArray(specifications) ? specifications : [specifications]).forEach(function(spec){

                self.$write_specification(f, base_path, self.$extract_resource_name(resource.path),spec);
            })
            return f.$write("\n\n");
        }, this);
        return this;
    };

    this.$write_specification = function (f, base_path, resource, specification) {
        var apis = specification["apis"];
        return apis && apis.map(function (method) {
            return method["operations"].map(function (operation) {
                return self.$write_operation(f, base_path, resource, operation, method["path"])
            })
        });
    };

    this.$write_operation = function (f, base_path, resource, operation, path) {
        var response;
        if (!operation.summary) {
            f.$write(this.$build_markdown_header("[Please add operation summary information to the summary section]\n\n", 3))
        } else {
            f.$write(this.$build_markdown_header(operation.summary + "\n", 3))
        }

        if (!operation.notes) {
            f.$write("[Please add operation information to the notes section]\n\n")
        } else {
            f.$write(operation.notes + "\n");
        }
        ;
        f.$write(this.$build_markdown_header("Definition", 4));
        f.$write("\n\n");
        this.$write_code_block(f, operation.httpMethod + " " + path);
        f.$write("\n\n");
        f.$write(this.$build_markdown_header("Arguments", 4));
        this.$write_arguments(f, operation.parameters);
        f.$write("\n\n");
        f.$write(this.$build_markdown_header("Example Request", 4));
        response = this.$write_example_request(f, base_path, operation, path, operation.parameters, resource);
        f.$write("\n\n");
        f.$write(this.$build_markdown_header("Example Response", 4));

        if (!response) {
            this.$write_code_block(f, response)
        }

        f.$write("\n\n");
        f.$write(this.$build_markdown_header("Potential Errors", 4));
        this.$write_errors(f, operation.errorResponses);
        return f.$write("\n\n");
    };

    this.$write_example_request = function (f, base_path, operation, path, arguments, resource) {

        path = this.$populate_arguments(path, arguments);
        var commmand, data;
        switch (operation.httpMethod) {
            case 'GET':
                command = "curl " + base_path + path;
                break;
            case 'POST':
            {

                data = resource ? "" : this.parameters[resource.toUpperCase() + ".POST"];
                command = "curl -X POST -H \"Content-Type:application/json\" -d '" + data + "' " + base_path + path;
                break;
            }
            case 'PUT':
            {
                data = resource ? "" : this.parameters[resource.toUpperCase() + ".PUT"];
                command = "curl -X PUT -H \"Content-Type:application/json\" -d '" + data + "' " + base_path + path;

            }
        }
        this.$write_code_block(f, command);
//        response = stdout.$read();
//        return (function () {
//            try {
//                __scope.JSON.$pretty_generate(__scope.JSON.$parse(response)).$gsub("\n", "\n    ")
//            } catch ($err) {
//                if (true) {
//                    response
//                }
//                else {
//                    throw $err;
//                }
//            }
//        }).call(this);
    };

    this.$populate_arguments = function (path, arguments) {
        path = path.replace("{format}", "json");
        if (!(arguments && arguments.length)) {
            return path;
        }

        arguments.filter(function (argument) {
            return argument.name && argument.paramType == 'path' && self.parameters && self.parameters[argument.name]
        }).map(function (argument) {
                return path = path.replace("{" + argument.name + "}", self.parameters[argument.name])
         });
        return path;
    };

    this.$write_errors = function (f, errors) {
        if (!(errors && errors.length)) {
            f.$write("* None\n");
            return null;
        }

        return errors.forEach(function (error) {

            f.$write("* ");
            if (!error.code) {
                f.$write("[Please add a code for error]")
            } else {
                f.$write("**" + error.code + "**");
            }

            if (!error.reason) {
                f.$write("")
            } else {
                f.$write(" - " + error.reason)
            }
            return f.$write("\n");
        });
    };

    this.$write_arguments = function (f, args) {
        if (!(args && args.length)) {
            f.$write("* None\n");
            return null;
        }

        return  args.forEach(function (argument) {
            f.$write("* ");
            if (!argument.name) {
                f.$write("[Please add a name for argument]")
            } else {
                f.$write("**" + argument.name + "**")
            }

            if (!argument.description) {
                f.$write("")
            } else {
                f.$write(" - " + argument.description)
            }

            return f.$write("\n");
        });
    };

    this.$write_code_block = function (f, text) {
        return f.$write("    " + text)
    };

    this.$write_header = function (f, api_name, api_version, base_path) {

        f.$write(this.$build_markdown_header(api_name + " " + api_version + " REST API", 1));
        f.$write("Base Path: " + base_path + "\n\n");
        f.$write(this.$build_input_here());
        f.$write("\n\n");
        f.$write(this.$build_markdown_header("General Considerations", 2));
        f.$write(this.$build_input_here());
        return f.$write("\n\n");
    };

    this.$extract_resource_name = function (path) {
        return path.substring(1, path.indexOf('.'));

    };

    this.$build_input_here = function () {
        return "[Please add API specific content here]\n"
    };

    this.$build_markdown_header = function (text, level) {
        var str = []
        while (level-- > 0)
            str.push('#');
        if (text)
            str.push(text);

        return str.join('') + "\n";
    };


};
module.exports = SwaggerToMarkdown
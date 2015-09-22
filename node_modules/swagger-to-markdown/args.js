var File = require('fs');
function opt(shrt, lng, help, type,  required) {
    return {
        short:shrt,
        long:lng,
        help:help,
        type:type,
        required:required
    }
}
function Parser(opts) {

    function help(err){
        var req = opts.filter(function(v){ return v.required}).map(function(v){ return '-'+ v.short+' '+ v.type }).join(' ');
        console.log('usage: '+process.argv.join(' '), req);
        if (!Array.isArray(err)){
            console.log('ERROR', Array.prototype.slice.call(arguments,0).join(' '));
            process.exit(1);
        }
        args.forEach(function(v){
            console.log(v.required ? '*': '', '-'+ v.short, '--'+ v.long, typeof v.type =='string' ? v.type : '', v.help)
        })
        process.exit(0)
    }
    this.handle = {
        'filename':function (args) {
            var file = args.shift();
            if (!File.existsSync(file)) {
                throw "file does not exist: " + file;
            }
            return file;
        },
        'json_file':function(args){
            var sargs =args.shift().split(/,/).map(function(v){
                    return JSON.parse(File.readFileSync(v, 'utf-8'))
                });

            return sargs.length == 1 ? sargs[0] : sargs;
        },
        'array':function (args) {
            return args.shift().split(/,\s*/)
        },
        'int':function (args) {
            return parseInt(args.shift())
        },
        'string':function(args){
            return args.shift();
        },
        'help':help
    }
    var self = this;
    function ex(opt,  args, ret){
        var f = typeof opt.type == 'function' ? opt.type : self.handle[opt.type];
        if (!f) {
            help('Could not handle opt', opt.long, opt);
        } else {
            try {
                ret[opt.long] = f.call(self, args);
            } catch (e) {
                help('Could not handle opt', e);
            }
        }
    }
    this.parse = function (args) {
        var ret = {};
        while (args.length) {
            var arg = args.shift();
            for (var i in opts) {
                var opt = opts[i]
                if ('-' + opt.short == arg || '--' + opt.long == arg) {
                    ex(opt, args, ret);
                }
            }
        }
        opts.forEach(function(opt){
            if (ret[opt.long])
                return;
            if (opt.default){
                ex(opt, [opt.default], ret);
            }else if (opt.required){
                help('required option "--'+ opt.long+ '" is not satisfied')
            }

        });
        return ret;
    }
}
Parser.opt = opt;
module.exports = Parser;
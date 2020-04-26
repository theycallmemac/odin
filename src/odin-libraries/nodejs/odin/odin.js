YAML = require('yamljs');
odinLogger = require('./odinLogger');

class Odin {
    constructor(config='job.yml', test=false) {
        this.config = YAML.load(config);
        this.test = test;
        this.id = this.config.job.id;
        // timestamp is used to identify approximate job run date time for observability
        this.timestamp = Date.now();
        this.logger = new odinLogger.OdinLogger();
        
        if (process.env.ODIN_EXEC_ENV) {
            this.ENV_CONFIG = true;
        }
        else {
            this.ENV_CONFIG = false;
        }
    }

    async condition(desc, expr){
        if (this.ENV_CONFIG) {
            this.logger.log('condition', desc, expr, this.id, this.timestamp);
        }
        return expr
    }

    async watch(desc, value){
         if (this.ENV_CONFIG) {
            this.logger.log('watch', desc, value, this.id, this.timestamp);
        }
    }

    async result(desc, status){
        if (this.ENV_CONFIG) {
            await this.logger.log('result',desc, status, this.id, this.timestamp);
        }
        process.exit(0);
    }
}

module.exports.Odin = Odin;
var fs = require('fs');
var os = require('os');
var Saxophone = require('saxophone');
var parser = new Saxophone();
var writer = fs.createWriteStream('node.csv');
var columns = [
    'AOGUID', 'FORMALNAME', 'REGIONCODE', 'AUTOCODE', 'AREACODE', 'CITYCODE', 'CTARCODE', 'PLACECODE', 'PLANCODE',
    'STREETCODE', 'EXTRCODE', 'SEXTCODE', 'OFFNAME', 'POSTALCODE', 'IFNSFL', 'TERRIFNSFL', 'IFNSUL', 'TERRIFNSUL',
    'OKATO', 'OKTMO', 'UPDATEDATE', 'SHORTNAME', 'AOLEVEL', 'PARENTGUID', 'AOID', 'PREVID', 'NEXTID', 'CODE',
    'PLAINCODE', 'ACTSTATUS', 'LIVESTATUS', 'CENTSTATUS', 'OPERSTATUS', 'CURRSTATUS', 'STARTDATE', 'ENDDATE', 'NORMDOC',
    'CADNUM', 'DIVTYPE',
];

var attrs = {};

parser
    .on('tagopen', node => {
        if (node.name === 'Object') {
            attrs = Saxophone.parseAttrs(node.attrs);

            writer.write(
                columns
                    .map(column => '"' + (attrs[column] || '').replace(/"/g, '""') + '"')
                    .join(',')
                + os.EOL,
            );
        }
    })
    .on('finish', () => {
        writer.close();
    });

// { highWaterMark: 64 * 1024 }
fs
    .createReadStream('./files/AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML')
    .pipe(parser);

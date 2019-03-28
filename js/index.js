var fs = require('fs');
var os = require('os');
var sax = require('sax').createStream();
var stream = fs.createWriteStream('node.csv');
var columns = [
    'AOGUID', 'FORMALNAME', 'REGIONCODE', 'AUTOCODE', 'AREACODE', 'CITYCODE', 'CTARCODE', 'PLACECODE', 'PLANCODE',
    'STREETCODE', 'EXTRCODE', 'SEXTCODE', 'OFFNAME', 'POSTALCODE', 'IFNSFL', 'TERRIFNSFL', 'IFNSUL', 'TERRIFNSUL',
    'OKATO', 'OKTMO', 'UPDATEDATE', 'SHORTNAME', 'AOLEVEL', 'PARENTGUID', 'AOID', 'PREVID', 'NEXTID', 'CODE',
    'PLAINCODE', 'ACTSTATUS', 'LIVESTATUS', 'CENTSTATUS', 'OPERSTATUS', 'CURRSTATUS', 'STARTDATE', 'ENDDATE', 'NORMDOC',
    'CADNUM', 'DIVTYPE',
];

sax
    .on('error', e => {
        // unhandled errors will throw, since this is a proper node
        // event emitter.
        console.error('error!', e);
        // clear the error
        sax._parser.error = null;
        sax._parser.resume();
    })
    .on('opentag', node => {
        if (node.name === 'OBJECT') {
            stream.write(
                columns
                    .map(column => '"' + (node.attributes[column] || '').replace(/"/g, '""') + '"')
                    .join(',')
                + os.EOL,
            );
        }
    })
    .on('end', () => {
        stream.close();
    });

// { highWaterMark: 64 * 1024 }
fs
    .createReadStream('../AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML')
    .pipe(sax);

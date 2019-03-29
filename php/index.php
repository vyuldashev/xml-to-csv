<?php

$columns = [
    'AOGUID',
    'FORMALNAME',
    'REGIONCODE',
    'AUTOCODE',
    'AREACODE',
    'CITYCODE',
    'CTARCODE',
    'PLACECODE',
    'PLANCODE',
    'STREETCODE',
    'EXTRCODE',
    'SEXTCODE',
    'OFFNAME',
    'POSTALCODE',
    'IFNSFL',
    'TERRIFNSFL',
    'IFNSUL',
    'TERRIFNSUL',
    'OKATO',
    'OKTMO',
    'UPDATEDATE',
    'SHORTNAME',
    'AOLEVEL',
    'PARENTGUID',
    'AOID',
    'PREVID',
    'NEXTID',
    'CODE',
    'PLAINCODE',
    'ACTSTATUS',
    'LIVESTATUS',
    'CENTSTATUS',
    'OPERSTATUS',
    'CURRSTATUS',
    'STARTDATE',
    'ENDDATE',
    'NORMDOC',
    'CADNUM',
    'DIVTYPE',
];

$result = fopen('php.csv', 'wb');

function startElement($parser, $name, $attrs)
{
    global $columns, $result;

    if ($name === 'OBJECT') {
        $addressObject = '';

        foreach ($columns as $column) {
            $addressObject .= ($attrs[$column] ?? '') . ';';
        }

        fwrite($result, $addressObject . PHP_EOL);
    }
}

function endElement($parser, $name)
{

}

$parser = xml_parser_create();

xml_set_element_handler($parser, 'startElement', 'endElement');

$f = fopen('./files/AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML', 'rb');

while ($data = fread($f, 32 * 1024)) {
    xml_parse($parser, $data);
}

xml_parser_free($parser);
fclose($f);
fclose($result);

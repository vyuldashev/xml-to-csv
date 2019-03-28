<?php
$start_time = microtime(true);

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

$xmlReader = new XMLReader();
$xmlReader->open('../AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML');

$result = fopen('php.csv', 'wb');

while ($xmlReader->read()) {
    if ($xmlReader->nodeType !== XMLReader::ELEMENT || $xmlReader->depth !== 1 || $xmlReader->localName !== 'Object') {
        continue;
    }

    $attributesCount = $xmlReader->attributeCount;

    $addressObject = [];

    for ($i = 0; $i < $attributesCount; $i++) {
        $xmlReader->moveToAttributeNo($i);

        if (in_array($xmlReader->name, $columns, true)) {
            $addressObject[$xmlReader->name] = $xmlReader->value;
        }
    }

    fputcsv($result, $addressObject);
}

echo microtime(true) - $start_time . PHP_EOL;

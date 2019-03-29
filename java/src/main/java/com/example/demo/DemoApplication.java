package com.example.demo;

import javax.xml.stream.XMLInputFactory;
import java.io.*;
import java.util.*;
import javax.xml.stream.FactoryConfigurationError;
import javax.xml.stream.XMLEventReader;
import javax.xml.stream.XMLStreamException;
import javax.xml.stream.events.Attribute;
import javax.xml.stream.events.StartElement;
import javax.xml.stream.events.XMLEvent;

public class DemoApplication {

    public static void main(String[] args) throws FileNotFoundException, XMLStreamException, FactoryConfigurationError {
        XMLInputFactory inputFactory = XMLInputFactory.newInstance();

        inputFactory.setProperty("javax.xml.stream.isCoalescing", true);

        // Setup a new eventReader
        InputStream in = new FileInputStream("./files/AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML");
        XMLEventReader eventReader = inputFactory.createXMLEventReader(in);

        String[] columns = {
                "AOGUID",
                "FORMALNAME",
                "REGIONCODE",
                "AUTOCODE",
                "AREACODE",
                "CITYCODE",
                "CTARCODE",
                "PLACECODE",
                "PLANCODE",
                "STREETCODE",
                "EXTRCODE",
                "SEXTCODE",
                "OFFNAME",
                "POSTALCODE",
                "IFNSFL",
                "TERRIFNSFL",
                "IFNSUL",
                "TERRIFNSUL",
                "OKATO",
                "OKTMO",
                "UPDATEDATE",
                "SHORTNAME",
                "AOLEVEL",
                "PARENTGUID",
                "AOID",
                "PREVID",
                "NEXTID",
                "CODE",
                "PLAINCODE",
                "ACTSTATUS",
                "LIVESTATUS",
                "CENTSTATUS",
                "OPERSTATUS",
                "CURRSTATUS",
                "STARTDATE",
                "ENDDATE",
                "NORMDOC",
                "CADNUM",
                "DIVTYPE",
        };

        File csvOutputFile = new File("ADDRESS_OBJECTS.csv");

        PrintWriter writer = new PrintWriter(csvOutputFile);

        while (eventReader.hasNext()) {
            XMLEvent event = eventReader.nextEvent();

            if (event.isStartElement()) {
                StartElement startElement = event.asStartElement();

                if (startElement.getName().getLocalPart().equals("Object")) {
                    HashMap<String, String> addressObject = new HashMap<>();

                    Iterator<Attribute> attributes = startElement.getAttributes();

                    while (attributes.hasNext()) {
                        Attribute attribute = attributes.next();

                        boolean contains = Arrays.asList(columns).contains(attribute.getName().toString());

                        if (contains) {
                            addressObject.put(attribute.getName().toString(), attribute.getValue());
                        }
                    }

                    // write to csv
                    String s = implode(";", addressObject);

                    writer.print(s);
                }
            }
        }
    }

    private static String implode(String delimiter, Map<String, String> map) {
        boolean first = true;
        StringBuilder sb = new StringBuilder();

        for (Map.Entry<String, String> e : map.entrySet()) {
            if (!first) {
                sb.append(delimiter);
            }

            sb.append(e.getValue());
            first = false;
        }

        sb.append("\n");

        return sb.toString();
    }
}

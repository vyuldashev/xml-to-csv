package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import javax.xml.stream.XMLInputFactory;
import java.io.*;
import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import javax.xml.stream.FactoryConfigurationError;
import javax.xml.stream.XMLEventReader;
import javax.xml.stream.XMLInputFactory;
import javax.xml.stream.XMLStreamException;
import javax.xml.stream.events.Attribute;
import javax.xml.stream.events.EndElement;
import javax.xml.stream.events.StartElement;
import javax.xml.stream.events.XMLEvent;


@SpringBootApplication
public class DemoApplication {

    public static void main(String[] args) throws FileNotFoundException, XMLStreamException, FactoryConfigurationError {
//        SpringApplication.run(DemoApplication.class, args);

        XMLInputFactory inputFactory = XMLInputFactory.newInstance();

        //inputFactory.setProperty("javax.xml.stream.isCoalescing", True)

        // Setup a new eventReader
        InputStream in = new FileInputStream("AS_ADDROBJ_20190324_a1a706ea-4ac7-43e7-b65b-68de81a57ddb.XML");
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

        PrintWriter pw = new PrintWriter(csvOutputFile);

        while (eventReader.hasNext()) {
            XMLEvent event = eventReader.nextEvent();

            HashMap<String, String> addressObject = new HashMap<>();

            //reach the start of an item
            if (event.isStartElement()) {

                StartElement startElement = event.asStartElement();

                if (startElement.getName().getLocalPart().equals("Object")) {

                    // attribute
                    Iterator<Attribute> attributes = startElement.getAttributes();
                    while (attributes.hasNext()) {
                        Attribute attribute = attributes.next();

                        boolean contains = Arrays.asList(columns).contains(attribute.getName().toString());

                        if (contains) {
                            addressObject.put(attribute.getName().toString(), attribute.getValue());
                        }
                    }
                }

                // data
                if (event.isStartElement()) {

//                    if (event.asStartElement().getName().getLocalPart().equals("thetext")) {
//                        event = eventReader.nextEvent();
//
//                        if (item.getFirstText() == null) {
//                            System.out.println("thetext: "
//                                    + event.asCharacters().getData());
//                            item.setFirstText("notnull");
//                            continue;
//                        } else {
//                            continue;
//                        }
//
//                    }
                }
            }

            //reach the end of an item
            if (event.isEndElement()) {
                EndElement endElement = event.asEndElement();
                if (endElement.getName().getLocalPart().equals("Object")) {
//                    item = null;
                }
            }

            // write to csv
            String s = implode(";", addressObject);

            pw.print(s);
        }

        System.out.println("JAva finished");
    }

    private static String implode(String delimiter, Map<String, String> map) {

        boolean first = true;
        StringBuilder sb = new StringBuilder();

        for (Map.Entry<String, String> e : map.entrySet()) {
            if (!first) sb.append(" " + delimiter + " ");
            sb.append(e.getValue());
            first = false;
        }

        return sb.toString();
    }
}

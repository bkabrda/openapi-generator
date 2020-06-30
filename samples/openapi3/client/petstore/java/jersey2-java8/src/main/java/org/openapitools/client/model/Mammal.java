/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import java.util.Objects;
import java.util.Arrays;
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.client.model.Pig;
import org.openapitools.client.model.Whale;
import org.openapitools.client.model.Zebra;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import org.openapitools.client.JSON;

import com.fasterxml.jackson.core.type.TypeReference;

import javax.ws.rs.core.GenericType;
import javax.ws.rs.core.Response;
import java.io.IOException;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.JsonToken;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.MapperFeature;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import org.openapitools.client.JSON;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
@JsonDeserialize(using = Mammal.MammalDeserializer.class)
@JsonSerialize(using = Mammal.MammalSerializer.class)
public class Mammal extends AbstractOpenApiSchema {
    private static final Logger log = Logger.getLogger(Mammal.class.getName());

    public static class MammalSerializer extends StdSerializer<Mammal> {
        public MammalSerializer(Class<Mammal> t) {
            super(t);
        }

        public MammalSerializer() {
            this(null);
        }

        @Override
        public void serialize(Mammal value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
            jgen.writeObject(value.getActualInstance());
        }
    }

    public static class MammalDeserializer extends StdDeserializer<Mammal> {
        public MammalDeserializer() {
            this(Mammal.class);
        }

        public MammalDeserializer(Class<?> vc) {
            super(vc);
        }

        @Override
        public Mammal deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
            JsonNode tree = jp.readValueAsTree();
            Object deserialized = null;
            Mammal newMammal = new Mammal();
            Map<String,Object> result2 = tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Map<String, Object>>() {});
            String discriminatorValue = (String)result2.get("className");
            switch (discriminatorValue) {
                case "Pig":
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Pig.class);
                    newMammal.setActualInstance(deserialized);
                    return newMammal;
                case "whale":
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Whale.class);
                    newMammal.setActualInstance(deserialized);
                    return newMammal;
                case "zebra":
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Zebra.class);
                    newMammal.setActualInstance(deserialized);
                    return newMammal;
                default:
                    log.log(Level.WARNING, String.format("Failed to lookup discriminator value `%s` for Mammal. Possible values: Pig whale zebra", discriminatorValue));
            }

            boolean typeCoercion = ctxt.isEnabled(MapperFeature.ALLOW_COERCION_OF_SCALARS);
            int match = 0;
            JsonToken token = tree.traverse(jp.getCodec()).nextToken();
            // deserialize Pig
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (Pig.class.equals(Integer.class) || Pig.class.equals(Long.class) || Pig.class.equals(Float.class) || Pig.class.equals(Double.class) || Pig.class.equals(Boolean.class) || Pig.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((Pig.class.equals(Integer.class) || Pig.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((Pig.class.equals(Float.class) || Pig.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (Pig.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (Pig.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Pig.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'Pig'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'Pig'", e);
            }

            // deserialize Whale
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (Whale.class.equals(Integer.class) || Whale.class.equals(Long.class) || Whale.class.equals(Float.class) || Whale.class.equals(Double.class) || Whale.class.equals(Boolean.class) || Whale.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((Whale.class.equals(Integer.class) || Whale.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((Whale.class.equals(Float.class) || Whale.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (Whale.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (Whale.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Whale.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'Whale'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'Whale'", e);
            }

            // deserialize Zebra
            try {
                boolean attemptParsing = true;
                // ensure that we respect type coercion as set on the client ObjectMapper
                if (Zebra.class.equals(Integer.class) || Zebra.class.equals(Long.class) || Zebra.class.equals(Float.class) || Zebra.class.equals(Double.class) || Zebra.class.equals(Boolean.class) || Zebra.class.equals(String.class)) {
                    attemptParsing = typeCoercion;
                    if (!attemptParsing) {
                        attemptParsing |= ((Zebra.class.equals(Integer.class) || Zebra.class.equals(Long.class)) && token == JsonToken.VALUE_NUMBER_INT);
                        attemptParsing |= ((Zebra.class.equals(Float.class) || Zebra.class.equals(Double.class)) && token == JsonToken.VALUE_NUMBER_FLOAT);
                        attemptParsing |= (Zebra.class.equals(Boolean.class) && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE));
                        attemptParsing |= (Zebra.class.equals(String.class) && token == JsonToken.VALUE_STRING);
                    }
                }
                if (attemptParsing) {
                    deserialized = tree.traverse(jp.getCodec()).readValueAs(Zebra.class);
                    // TODO: there is no validation against JSON schema constraints
                    // (min, max, enum, pattern...), this does not perform a strict JSON
                    // validation, which means the 'match' count may be higher than it should be.
                    match++;
                    log.log(Level.FINER, "Input data matches schema 'Zebra'");
                }
            } catch (Exception e) {
                // deserialization failed, continue
                log.log(Level.FINER, "Input data does not match schema 'Zebra'", e);
            }

            if (match == 1) {
                Mammal ret = new Mammal();
                ret.setActualInstance(deserialized);
                return ret;
            }
            throw new IOException(String.format("Failed deserialization for Mammal: %d classes match result, expected 1", match));
        }

        /**
         * Handle deserialization of the 'null' value.
         */
        @Override
        public Mammal getNullValue(DeserializationContext ctxt) throws JsonMappingException {
            throw new JsonMappingException(ctxt.getParser(), "Mammal cannot be null");
        }
    }

    // store a list of schema names defined in oneOf
    public final static Map<String, GenericType> schemas = new HashMap<String, GenericType>();

    public Mammal() {
        super("oneOf", Boolean.FALSE);
    }
  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   */
  @JsonAnySetter
  public Mammal putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   */
  @JsonAnyGetter
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }

    /**
     * Return true if this mammal object is equal to o.
     */
    @Override
    public boolean equals(Object o) {
        return super.equals(o) && Objects.equals(this.additionalProperties, ((Mammal)o).additionalProperties);
    }

    @Override
    public int hashCode() {
        return Objects.hash(getActualInstance(), isNullable(), getSchemaType(), additionalProperties);
    }
    public Mammal(Pig o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public Mammal(Whale o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public Mammal(Zebra o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    static {
        schemas.put("Pig", new GenericType<Pig>() {
        });
        schemas.put("Whale", new GenericType<Whale>() {
        });
        schemas.put("Zebra", new GenericType<Zebra>() {
        });
        JSON.registerDescendants(Mammal.class, Collections.unmodifiableMap(schemas));
        // Initialize and register the discriminator mappings.
        Map<String, Class<?>> mappings = new HashMap<String, Class<?>>();
        mappings.put("Pig", Pig.class);
        mappings.put("whale", Whale.class);
        mappings.put("zebra", Zebra.class);
        mappings.put("mammal", Mammal.class);
        JSON.registerDiscriminator(Mammal.class, "className", mappings);
    }

    @Override
    public Map<String, GenericType> getSchemas() {
        return Mammal.schemas;
    }

    /**
     * Set the instance that matches the oneOf child schema, check
     * the instance parameter is valid against the oneOf child schemas.
     *
     * It could be an instance of the 'oneOf' schemas.
     * The oneOf child schemas may themselves be a composed schema (allOf, anyOf, oneOf).
     */
    @Override
    public void setActualInstance(Object instance) {
        if (JSON.isInstanceOf(Pig.class, instance, new HashSet<Class<?>>())) {
            super.setActualInstance(instance);
            return;
        }

        if (JSON.isInstanceOf(Whale.class, instance, new HashSet<Class<?>>())) {
            super.setActualInstance(instance);
            return;
        }

        if (JSON.isInstanceOf(Zebra.class, instance, new HashSet<Class<?>>())) {
            super.setActualInstance(instance);
            return;
        }

        throw new RuntimeException("Invalid instance type. Must be Pig, Whale, Zebra");
    }



}


/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI-Generator 4.3.0-SNAPSHOT.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



#include "Pet.h"

namespace org {
namespace openapitools {
namespace client {
namespace model {




Pet::Pet()
{
    m_Id = 0L;
    m_IdIsSet = false;
    m_CategoryIsSet = false;
    m_Name = utility::conversions::to_string_t("");
    m_TagsIsSet = false;
    m_Status = utility::conversions::to_string_t("");
    m_StatusIsSet = false;
}

Pet::~Pet()
{
}

void Pet::validate()
{
    // TODO: implement validation
}

web::json::value Pet::toJson() const
{
    web::json::value val = web::json::value::object();

    if(m_IdIsSet)
    {
        val[utility::conversions::to_string_t("id")] = ModelBase::toJson(m_Id);
    }
    if(m_CategoryIsSet)
    {
        val[utility::conversions::to_string_t("category")] = ModelBase::toJson(m_Category);
    }
    val[utility::conversions::to_string_t("name")] = ModelBase::toJson(m_Name);
    {
        std::vector<web::json::value> jsonArray;
        for( auto& item : m_PhotoUrls )
        {
            jsonArray.push_back(ModelBase::toJson(item));
        }
        val[utility::conversions::to_string_t("photoUrls")] = web::json::value::array(jsonArray);
    }
    {
        std::vector<web::json::value> jsonArray;
        for( auto& item : m_Tags )
        {
            jsonArray.push_back(ModelBase::toJson(item));
        }
        if(jsonArray.size() > 0)
        {
            val[utility::conversions::to_string_t("tags")] = web::json::value::array(jsonArray);
        }
    }
    if(m_StatusIsSet)
    {
        val[utility::conversions::to_string_t("status")] = ModelBase::toJson(m_Status);
    }

    return val;
}

void Pet::fromJson(const web::json::value& val)
{
    if(val.has_field(utility::conversions::to_string_t("id")))
    {
        const web::json::value& fieldValue = val.at(utility::conversions::to_string_t("id"));
        if(!fieldValue.is_null())
        {
            setId(ModelBase::int64_tFromJson(fieldValue));
        }
    }
    if(val.has_field(utility::conversions::to_string_t("category")))
    {
        const web::json::value& fieldValue = val.at(utility::conversions::to_string_t("category"));
        if(!fieldValue.is_null())
        {
            auto newItem = std::make_shared<std::shared_ptr<Category>::element_type>();
            newItem->fromJson(fieldValue);
            setCategory( newItem );
        }
    }
    setName(ModelBase::stringFromJson(val.at(utility::conversions::to_string_t("name"))));
    {
        m_PhotoUrls.clear();
        std::vector<web::json::value> jsonArray;
        for( auto& item : val.at(utility::conversions::to_string_t("photoUrls")).as_array() )
        {
            m_PhotoUrls.push_back(ModelBase::stringFromJson(item));
        }
    }
    {
        m_Tags.clear();
        std::vector<web::json::value> jsonArray;
        if(val.has_field(utility::conversions::to_string_t("tags")))
        {
        for( auto& item : val.at(utility::conversions::to_string_t("tags")).as_array() )
        {
            if(item.is_null())
            {
                m_Tags.push_back( std::shared_ptr<Tag>(nullptr) );
            }
            else
            {
                auto newItem = std::make_shared<std::shared_ptr<Tag>::element_type>();
                newItem->fromJson(item);
                m_Tags.push_back( newItem );
            }
        }
        }
    }
    if(val.has_field(utility::conversions::to_string_t("status")))
    {
        const web::json::value& fieldValue = val.at(utility::conversions::to_string_t("status"));
        if(!fieldValue.is_null())
        {
            setStatus(ModelBase::stringFromJson(fieldValue));
        }
    }
}

void Pet::toMultipart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& prefix) const
{
    utility::string_t namePrefix = prefix;
    if(namePrefix.size() > 0 && namePrefix.substr(namePrefix.size() - 1) != utility::conversions::to_string_t("."))
    {
        namePrefix += utility::conversions::to_string_t(".");
    }

    if(m_IdIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + utility::conversions::to_string_t("id"), m_Id));
    }
    if(m_CategoryIsSet)
    {
        if (m_Category.get())
        {
            m_Category->toMultipart(multipart, utility::conversions::to_string_t("category."));
        }
    }
    multipart->add(ModelBase::toHttpContent(namePrefix + utility::conversions::to_string_t("name"), m_Name));
    {
        std::vector<web::json::value> jsonArray;
        for( auto& item : m_PhotoUrls )
        {
            jsonArray.push_back(ModelBase::toJson(item));
        }
        multipart->add(ModelBase::toHttpContent(namePrefix + utility::conversions::to_string_t("photoUrls"), web::json::value::array(jsonArray), utility::conversions::to_string_t("application/json")));
            }
    {
        std::vector<web::json::value> jsonArray;
        for( auto& item : m_Tags )
        {
            jsonArray.push_back(ModelBase::toJson(item));
        }
        
        if(jsonArray.size() > 0)
        {
            multipart->add(ModelBase::toHttpContent(namePrefix + utility::conversions::to_string_t("tags"), web::json::value::array(jsonArray), utility::conversions::to_string_t("application/json")));
        }
    }
    if(m_StatusIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + utility::conversions::to_string_t("status"), m_Status));
    }
}

void Pet::fromMultiPart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& prefix)
{
    utility::string_t namePrefix = prefix;
    if(namePrefix.size() > 0 && namePrefix.substr(namePrefix.size() - 1) != utility::conversions::to_string_t("."))
    {
        namePrefix += utility::conversions::to_string_t(".");
    }

    if(multipart->hasContent(utility::conversions::to_string_t("id")))
    {
        setId(ModelBase::int64_tFromHttpContent(multipart->getContent(utility::conversions::to_string_t("id"))));
    }
    if(multipart->hasContent(utility::conversions::to_string_t("category")))
    {
        if(multipart->hasContent(utility::conversions::to_string_t("category")))
        {
            auto newItem = std::make_shared<std::shared_ptr<Category>::element_type>();
            newItem->fromMultiPart(multipart, utility::conversions::to_string_t("category."));
            setCategory( newItem );
        }
    }
    setName(ModelBase::stringFromHttpContent(multipart->getContent(utility::conversions::to_string_t("name"))));
    {
        m_PhotoUrls.clear();

        web::json::value jsonArray = web::json::value::parse(ModelBase::stringFromHttpContent(multipart->getContent(utility::conversions::to_string_t("photoUrls"))));
        for( auto& item : jsonArray.as_array() )
        {
            m_PhotoUrls.push_back(ModelBase::stringFromJson(item));
        }
    }
    {
        m_Tags.clear();
        if(multipart->hasContent(utility::conversions::to_string_t("tags")))
        {

        web::json::value jsonArray = web::json::value::parse(ModelBase::stringFromHttpContent(multipart->getContent(utility::conversions::to_string_t("tags"))));
        for( auto& item : jsonArray.as_array() )
        {
            if(item.is_null())
            {
                m_Tags.push_back( std::shared_ptr<Tag>(nullptr) );
            }
            else
            {
                auto newItem = std::make_shared<std::shared_ptr<Tag>::element_type>();
                newItem->fromJson(item);
                m_Tags.push_back( newItem );
            }
        }
        }
    }
    if(multipart->hasContent(utility::conversions::to_string_t("status")))
    {
        setStatus(ModelBase::stringFromHttpContent(multipart->getContent(utility::conversions::to_string_t("status"))));
    }
}

int64_t Pet::getId() const
{
    return m_Id;
}

void Pet::setId(int64_t value)
{
    m_Id = value;
    m_IdIsSet = true;
}

bool Pet::idIsSet() const
{
    return m_IdIsSet;
}

void Pet::unsetId()
{
    m_IdIsSet = false;
}

std::shared_ptr<Category> Pet::getCategory() const
{
    return m_Category;
}

void Pet::setCategory(const std::shared_ptr<Category>& value)
{
    m_Category = value;
    m_CategoryIsSet = true;
}

bool Pet::categoryIsSet() const
{
    return m_CategoryIsSet;
}

void Pet::unsetCategory()
{
    m_CategoryIsSet = false;
}

utility::string_t Pet::getName() const
{
    return m_Name;
}

void Pet::setName(const utility::string_t& value)
{
    m_Name = value;
    
}

std::vector<utility::string_t>& Pet::getPhotoUrls()
{
    return m_PhotoUrls;
}

void Pet::setPhotoUrls(const std::vector<utility::string_t>& value)
{
    m_PhotoUrls = value;
    
}

std::vector<std::shared_ptr<Tag>>& Pet::getTags()
{
    return m_Tags;
}

void Pet::setTags(const std::vector<std::shared_ptr<Tag>>& value)
{
    m_Tags = value;
    m_TagsIsSet = true;
}

bool Pet::tagsIsSet() const
{
    return m_TagsIsSet;
}

void Pet::unsetTags()
{
    m_TagsIsSet = false;
}

utility::string_t Pet::getStatus() const
{
    return m_Status;
}

void Pet::setStatus(const utility::string_t& value)
{
    m_Status = value;
    m_StatusIsSet = true;
}

bool Pet::statusIsSet() const
{
    return m_StatusIsSet;
}

void Pet::unsetStatus()
{
    m_StatusIsSet = false;
}

}
}
}
}



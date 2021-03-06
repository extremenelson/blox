/*
 * Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"). You may
 * not use this file except in compliance with the License. A copy of the
 * License is located at
 *
 *     http://aws.amazon.com/apache2.0/
 *
 * or in the "LICENSE" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */
package com.amazonaws.blox.frontend.operations;

import com.amazonaws.blox.dataservicemodel.v1.exception.ClientException;
import com.amazonaws.blox.dataservicemodel.v1.exception.ServiceException;
import com.amazonaws.blox.frontend.mappers.DeleteEnvironmentMapper;
import com.amazonaws.blox.frontend.models.Environment;
import io.swagger.annotations.ApiOperation;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DeleteEnvironment extends EnvironmentController {

  @Autowired DeleteEnvironmentMapper mapper;

  @RequestMapping(path = "/{environmentName}", method = RequestMethod.DELETE)
  @ApiOperation(value = "Delete an Environment by name")
  public DeleteEnvironmentResponse deleteEnvironment(
      @PathVariable("cluster") String cluster,
      @PathVariable("environmentName") String environmentName,
      @RequestParam(value = "forceDelete", required = false, defaultValue = "false")
          boolean forceDelete)
      throws ServiceException, ClientException {

    return mapper.fromDataServiceResponse(
        dataService.deleteEnvironment(
            mapper.toDataServiceRequest(
                getApiGatewayRequestContext(), cluster, environmentName, forceDelete)));
  }

  @Data
  @Builder
  @AllArgsConstructor
  @NoArgsConstructor
  public static class DeleteEnvironmentResponse {

    private Environment environment;
  }
}

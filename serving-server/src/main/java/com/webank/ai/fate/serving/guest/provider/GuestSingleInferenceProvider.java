package com.webank.ai.fate.serving.guest.provider;

import com.google.common.collect.Maps;
import com.google.common.util.concurrent.ListenableFuture;
import com.webank.ai.fate.serving.core.bean.*;
import com.webank.ai.fate.serving.core.model.Model;
import com.webank.ai.fate.serving.core.model.ModelProcessor;
import com.webank.ai.fate.serving.core.rpc.core.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;
import java.util.concurrent.Future;


@FateService(name = "singleInference", preChain = {
        "monitorInterceptor",
        "requestOverloadBreaker",
        "guestSingleParamInterceptor",
        "guestModelInterceptor",
        "modelOverloadBreaker",
        "federationRouterInterceptor"
}, postChain = {
        "monitorInterceptor"
})
@Service
public class GuestSingleInferenceProvider extends AbstractServingServiceProvider<InferenceRequest, ReturnResult> {
    @Autowired
    FederatedRpcInvoker federatedRpcInvoker;

    @Value("${inference.single.timeout:3000}")
    long timeout;

    @Autowired
    Environment environment;

    @Override
    public ReturnResult doService(Context context, InboundPackage inboundPackage, OutboundPackage outboundPackage) {
        Model model = ((ServingServerContext) context).getModel();
        ModelProcessor modelProcessor = model.getModelProcessor();
        InferenceRequest inferenceRequest = (InferenceRequest) inboundPackage.getBody();
        Map<String, Future> futureMap = Maps.newHashMap();
        modelProcessor.guestPrepareDataBeforeInference(context, inferenceRequest);
        List<FederatedRpcInvoker.RpcDataWraper> rpcList = this.buildRpcDataWraper(context, Dict.FEDERATED_INFERENCE, inferenceRequest);

        Boolean useCache = environment.getProperty(Dict.PROPERTY_REMOTE_MODEL_INFERENCE_RESULT_CACHE_SWITCH, boolean.class, true);

        rpcList.forEach((rpcDataWraper -> {
            // ListenableFuture<Proxy.Packet> future = federatedRpcInvoker.async(context, rpcDataWraper);
            ListenableFuture<ReturnResult> future = federatedRpcInvoker.singleInferenceRpcWithCache(context, rpcDataWraper, useCache);

            futureMap.put(rpcDataWraper.getHostModel().getPartId(), future);
        }));
        ReturnResult returnResult = modelProcessor.guestInference(context, inferenceRequest, futureMap, timeout);
        return returnResult;
    }


    @Override
    protected OutboundPackage<ReturnResult> serviceFailInner(Context context, InboundPackage<InferenceRequest> data, Throwable e) {
        OutboundPackage<ReturnResult> outboundPackage = new OutboundPackage<ReturnResult>();
        ReturnResult returnResult = ErrorMessageUtil.handleExceptionToReturnResult(e);
        outboundPackage.setData(returnResult);
        context.setReturnCode(returnResult.getRetcode());
        return outboundPackage;
    }
}
